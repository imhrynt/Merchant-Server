package balance

import (
	"net/http"
	"server/model"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store    model.BalanceInterface
	validate *validator.Validate
}

func NewHandler(s model.BalanceInterface, v *validator.Validate) *Handler {
	return &Handler{store: s, validate: v}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/balance/build", h.handleBuild).Methods("POST")
	router.HandleFunc("/balance/read", h.handleFetch).Methods("GET")
	router.HandleFunc("/balance/", h.handleModification).Methods("POST")
}

func (h *Handler) handleBuild(writer http.ResponseWriter, request *http.Request)        {}
func (h *Handler) handleFetch(writer http.ResponseWriter, request *http.Request)        {}
func (h *Handler) handleModification(writer http.ResponseWriter, request *http.Request) {}
