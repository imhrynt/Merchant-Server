package model

import (
	"context"
	"net/http"
)

type UserInterface interface {
	Authentication(w http.ResponseWriter, ctx context.Context, user UserAuthenticateRequest)
	Register(w http.ResponseWriter, ctx context.Context, user UserRegisterRequest)
	Fetch(w http.ResponseWriter, ctx context.Context)
	Modification(w http.ResponseWriter, ctx context.Context, user UserModifyRequest)
}

type BalanceInterface interface {
	//Create(ctx context.Context, balance BalanceCreateRequest) int64
	//ReadByUser(ctx context.Context) BalanceResponse
	//Add(ctx context.Context, balance BalanceCreateRequest) int64
	//Sub(ctx context.Context, balance BalanceCreateRequest) int64
}

type RoleInterface interface {
	Create(ctx context.Context, role RoleCreateRequest) int64
	Read(ctx context.Context) []RoleResponse
	Update(ctx context.Context, role RoleUpdateRequest) int64
}

type TransactionInterface interface {
	Create(ctx context.Context, transaction TransactionCreateRequest) int64
	Read(ctx context.Context) []TransactionResponse
	//ReadByUser(ctx context.Context, user U) []TransactionResponse
	Update(ctx context.Context, transaction TransactionUpdateRequest) int64
}
