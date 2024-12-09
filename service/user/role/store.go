package role

import (
	"context"
	"database/sql"
	"log"
	"server/model"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) Create(ctx context.Context, role model.RoleCreateRequest) int64 {
	SQL := "INSERT INTO roles (role) VALUES (?)"
	result, err := s.db.ExecContext(ctx, SQL, role.Role)
	if err != nil {
		log.Fatalln(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}
	return id
}

func (s *Store) Read(ctx context.Context) []model.RoleResponse {
	SQL := "SELECT id, role FROM roles"
	rows, err := s.db.QueryContext(ctx, SQL)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	var roles []model.RoleResponse
	for rows.Next() {
		role := model.RoleResponse{}
		if err := rows.Scan(
			&role.Id,
			&role.Role,
		); err != nil {
			log.Fatalln(err)
		}
		roles = append(roles, role)
	}

	return roles
}

func (s *Store) Update(ctx context.Context, role model.RoleUpdateRequest) int64 {
	SQL := "UPDATE role SET role = ?, update_at = now() WHERE id = ?"
	result, err := s.db.ExecContext(ctx, SQL, role.Role, role.Id)
	if err != nil {
		log.Fatalln(err)
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}

	return rowAffected
}

func (s *Store) ReadById(ctx context.Context, id int64) model.RoleResponse {
	return model.RoleResponse{}
}
