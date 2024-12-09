package role

import (
	"net/http"
	"server/model"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store    model.RoleInterface
	validate *validator.Validate
}

func NewHandler(s model.RoleInterface, v *validator.Validate) *Handler {
	return &Handler{store: s, validate: v}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/role/build", h.handleBuild).Methods("POST")
	router.HandleFunc("/role/fetch", h.handleFetch).Methods("GET")
	router.HandleFunc("/role/modification", h.handleModification).Methods("POST")
}

func (h *Handler) handleBuild(writer http.ResponseWriter, request *http.Request)        {}
func (h *Handler) handleFetch(writer http.ResponseWriter, request *http.Request)        {}
func (h *Handler) handleModification(writer http.ResponseWriter, request *http.Request) {}
