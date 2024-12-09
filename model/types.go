package model

import "time"

/*
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
*/

/*
/////////////////////////
	RESPONSE STRUCT
/////////////////////////
*/

type Response struct {
	Code   int64       `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  string      `json:"message,omitempty"`
}

/*
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
*/

/*
/////////////////////////
	USERS STRUCT
/////////////////////////
*/

type UserAuthenticateRequest struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}

type UserRegisterRequest struct {
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required" json:"email"`
	Phone    string `validate:"required" json:"phone"`
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}

type UserModifyRequest struct {
	Id       int64  `validate:"required" json:"id"`
	Role_id  int64  `validate:"required" json:"role_id"`
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required" json:"email"`
	Phone    string `validate:"required" json:"phone"`
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
	Img_URL  string `validate:"required" json:"img_url"`
}

type UserResponse struct {
	Id       int64          `json:"id"`
	Role     string         `json:"role"`
	Wallet   WalletResponse `json:"walle,omitemptyt"`
	Name     string         `json:"name"`
	Email    string         `json:"email"`
	Phone    string         `json:"phone"`
	Username string         `json:"username"`
	Password string         `json:"-"`
	Img_URL  *string        `json:"img_url,omitempty"`
}

/*
/////////////////////////
	BALANCE STRUCT
/////////////////////////
*/

type WalletCreateRequest struct {
	User_id int64 `validate:"required" json:"user_id"`
	Balance int64 `validate:"required" json:"balance"`
}

type WalletUpdateRequest struct {
	Id      int64 `validate:"required" json:"id"`
	User_id int64 `validate:"required" json:"user_id"`
	Balance int64 `validate:"required" json:"balance"`
}

type WalletResponse struct {
	Id      int64 `json:"id"`
	Balance int64 `json:"balance"`
}

/*
/////////////////////////
	ROLE STRUCT
/////////////////////////
*/

type RoleCreateRequest struct {
	Role string `validate:"required" json:"role"`
}

type RoleUpdateRequest struct {
	Id   int64  `validate:"required" json:"id"`
	Role string `validate:"required" json:"role"`
}

type RoleResponse struct {
	Id   int64  `json:"id"`
	Role string `json:"role"`
}

/*
/////////////////////////
	TRANSACTION STRUCT
/////////////////////////
*/

type TransactionCreateRequest struct {
	User_id        int64     `validate:"required" json:"user_id"`
	Provider_id    int64     `validate:"required" json:"provider_id"`
	Payment_method string    `validate:"required" json:"payment_method"`
	Payment_name   string    `validate:"required" json:"payment_name"`
	Total          int64     `validate:"required" json:"total"`
	Expired_at     time.Time `validate:"required" json:"expired_at"`
}

type TransactionUpdateRequest struct {
	Id         int64     `validate:"required" json:"id"`
	Status     string    `validate:"required" json:"status"`
	Payment_at time.Time `validate:"required" json:"payment_at"`
}

type TransactionResponse struct {
	Id         int64     `json:"id"`
	Payment_at time.Time `json:"payment_at"`
	Expired_at time.Time `json:"expired_at"`
	Created_at time.Time `json:"created_at"`
}

/*
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
*/

/*
/////////////////////////
	PROVIDER STRUCT
/////////////////////////
*/

type Provider struct {
	Id            int64     `json:"id"`
	Provider      string    `json:"provider"`
	Api_key       string    `json:""`
	Private_key   string    `json:""`
	Merchant_code string    `json:""`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
}

/*
/////////////////////////
	PAYMENT STRUCT
/////////////////////////
*/

type Payment struct {
	Id          int64     `json:"id"`
	User_id     int64     `json:"user_id"`
	Provider_id int64     `json:"provider_id"`
	Total_cost  int64     `json:"total_cost"`
	Payment_url string    `json:"payment_url"`
	Status      string    `json:"status"`
	Payment_at  time.Time `json:"payment_at"`
	Expired_at  time.Time `json:"expired_at"`
	Created_at  time.Time `json:"created_at"`
}

/*
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
*/

//////////////////////////////////////////////
//////////		CATEGORY MODULE
//////////////////////////////////////////////

type Category struct {
	Id         int64     `json:"id"`
	Category   string    `json:"category"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
