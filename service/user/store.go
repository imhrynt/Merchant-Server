package user

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"server/model"
	"server/util"

	"golang.org/x/crypto/bcrypt"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

var userExist = errors.New("Account already exists")

// Authentication
func (s *Store) Authentication(w http.ResponseWriter, ctx context.Context, v model.UserAuthenticateRequest) {
	var user model.UserResponse
	SQL := "SELECT u.id, r.role, u.name, u.email, u.phone, u.username, u.password, u.img_url FROM users u INNER JOIN roles r ON u.role_id = r.id WHERE u.username = ?"
	if err := s.db.QueryRowContext(ctx, SQL, v.Username).Scan(&user.Id, &user.Role, &user.Name, &user.Email, &user.Phone, &user.Username, &user.Password, &user.Img_URL); err != nil {
		if err == sql.ErrNoRows {
			util.WriteJSON(w, 401, &model.Response{
				Code:   401,
				Status: "UNAUTHORIZED",
				Error:  "Authorization failed"})
		} else {
			util.WriteJSON(w, 500, &model.Response{
				Code:   500,
				Status: "INTERNAL_ERROR",
				Error:  "Error fetching user data"})
		}
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(v.Password)); err != nil {
		util.WriteJSON(w, 401, &model.Response{
			Code:   401,
			Status: "UNAUTHORIZED",
			Error:  "Authorization failed"})
		return
	}
	util.WriteJSON(w, 200, &model.Response{
		Code:   200,
		Status: "OK",
		Data:   user})
}

// Register
func (s *Store) Register(w http.ResponseWriter, ctx context.Context, user model.UserRegisterRequest) {
	if err := s.accountExist(ctx, user); err != nil {
		if err == userExist {
			util.WriteJSON(w, 409, &model.Response{
				Code:   409,
				Status: "CONFLICT",
				Error:  userExist.Error()})
		} else {
			util.WriteJSON(w, 500, &model.Response{
				Code:   500,
				Status: "INTERNAL_ERROR",
				Error:  "Failed to check account existence"})
		}
		return
	}
	password, err := util.ENCRYPT_BCRYPT([]byte(user.Password), 12)
	if err != nil {
		util.WriteJSON(w, 500, &model.Response{
			Code:   500,
			Status: "INTERNAL_ERROR",
			Error:  "Failed to hash password"})
		return
	}
	SQL := "INSERT INTO users (name, email, phone, username, password) VALUES (?, ?, ?, ?, ?)"
	result, err := s.db.ExecContext(ctx, SQL, user.Name, user.Email, user.Phone, user.Username, password)
	if err != nil {
		util.WriteJSON(w, 500, &model.Response{
			Code:   500,
			Status: "INTERNAL_ERROR",
			Error:  "Failed to insert user"})
		return
	}
	id, err := result.LastInsertId()
	if err != nil || id > 0 {
		util.WriteJSON(w, 500, &model.Response{
			Code:   500,
			Status: "INTERNAL_ERROR",
			Error:  "Failed to retrieve last insert ID"})
		return
	}
	SQL = "INSERT INTO balances (user_id) VALUES (?)"
	result, err = s.db.ExecContext(ctx, SQL, id)
	if err != nil {
		util.WriteJSON(w, 500, &model.Response{
			Code:   500,
			Status: "INTERNAL_ERROR",
			Error:  "Failed to insert user"})
		return
	}
	util.WriteJSON(w, 201, &model.Response{
		Code:   201,
		Status: "CREATED",
		Data:   map[string]interface{}{"user_id": id}})
}

func (s *Store) Fetch(writer http.ResponseWriter, ctx context.Context) { // VALID
	users, response := []model.UserResponse{}, &model.Response{}
	SQL := "SELECT u.id, u.role_id, b.id, b.balance, u.name, u.email, u.phone, u.username, u.img_url FROM users u INNER JOIN balances b ON u.id = b.user_id"
	rows, err := s.db.QueryContext(ctx, SQL)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Status = "INTERNAL_ERROR"
		response.Error = "Failed to fetch users"
		util.WriteJSON(writer, http.StatusInternalServerError, response)
		return
	}
	defer rows.Close()
	for rows.Next() {
		user := model.UserResponse{}
		var image_url *string
		if err := rows.Scan(
			&user.Id,
			&user.Role,
			&user.Wallet.Id,
			&user.Wallet.Balance,
			&user.Name,
			&user.Email,
			&user.Phone,
			&user.Username,
			&image_url,
		); err != nil {
			response.Code = http.StatusInternalServerError
			response.Status = "INTERNAL_ERROR"
			response.Error = "Failed to process user data"
			util.WriteJSON(writer, http.StatusInternalServerError, response)
			return
		}
		users = append(users, user)
	}
	if len(users) == 0 {
		response.Code = http.StatusNotFound
		response.Status = "NOT FOUND"
		response.Error = "No users found"
		util.WriteJSON(writer, http.StatusNotFound, response)
		return
	}
	response.Code = http.StatusOK
	response.Status = "OK"
	response.Data = users
	util.WriteJSON(writer, http.StatusOK, response)
}

func (s *Store) Modification(writer http.ResponseWriter, ctx context.Context, user model.UserModifyRequest) {
	response := &model.Response{}
	var params []interface{}
	SQL := "UPDATE users SET"
	if user.Name != "" {
		params = append(params, user.Name)
		SQL += " name = ?,"
	}
	if user.Email != "" {
		params = append(params, user.Email)
		SQL += " email = ?,"
	}
	if user.Phone != "" {
		params = append(params, user.Phone)
		SQL += " phone = ?,"
	}
	if user.Password != "" {
		params = append(params, user.Password)
		SQL += " password = ?,"
	}
	if user.Img_URL != "" {
		params = append(params, user.Img_URL)
		SQL += " img_url = ?,"
	}
	params = append(params, user.Id)
	result, err := s.db.ExecContext(ctx, SQL[:len(SQL)-1]+" WHERE id = ?", params...)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Status = "ERROR"
		response.Error = "failed to update user"
		util.WriteJSON(writer, http.StatusInternalServerError, response)
		return
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		response.Code = http.StatusOK
		response.Status = "ERROR"
		response.Error = err.Error()
		util.WriteJSON(writer, http.StatusOK, response)
		return
	}
	response.Code = http.StatusOK
	response.Status = "OK"
	response.Data = map[string]interface{}{"row_affected": rowAffected}
	util.WriteJSON(writer, http.StatusOK, response)
}

func basicCheck(user model.UserRegisterRequest) (int, string, string) {
	var message string
	if user.Name == "" {
		message = "Missing fullname field"
	} else if user.Email == "" {
		message = "Missing email field"
	} else if user.Password == "" {
		message = "Missing email field"
	} else if user.Username == "" {
		message = "Missing email field"
	} else if user.Password == "" {
		message = "Missing email field"
	}
	return http.StatusBadRequest, "BAD_REQUEST", message
}

func (s *Store) accountExist(ctx context.Context, user model.UserRegisterRequest) error {
	SQL := "SELECT * FROM users WHERE email = ? OR phone = ? OR username = ? LIMIT 1"
	var exist int
	if err := s.db.QueryRowContext(ctx, SQL, user.Email, user.Phone, user.Username).Scan(&exist); err != nil {
		return err
	} else if exist > 0 {
		return userExist
	}
	return nil
}
