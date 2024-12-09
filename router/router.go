package router

import (
	"database/sql"
	"net/http"
	"server/service/user"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type API struct {
	port string
	db   *sql.DB
}

func NewAPI(port string, db *sql.DB) *API {
	return &API{
		port: port,
		db:   db,
	}
}

func (a *API) Run() error {
	router := mux.NewRouter()
	validator := validator.New()
	subrouter := router.PathPrefix("/api/v1/services").Subrouter()
	userStore := user.NewStore(a.db)
	userHandler := user.NewHandler(userStore, validator)
	userHandler.Routing(subrouter)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))
	return http.ListenAndServe(a.port, router)
}
