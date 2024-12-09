package user

import (
	"net/http"
	"server/model"
	"server/util"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store    model.UserInterface
	validate *validator.Validate
}

func NewHandler(s model.UserInterface, v *validator.Validate) *Handler {
	return &Handler{store: s, validate: v}
}

func (h *Handler) Routing(r *mux.Router) {
	r.HandleFunc("/user/authentication", h.handleAuthentication).Methods("POST")
	r.HandleFunc("/user/register", h.Register).Methods("POST")
	r.HandleFunc("/user/fetch", h.Fetch).Methods("GET")
	r.HandleFunc("/user/modification", h.Modification).Methods("POST")
}

func (h *Handler) handleAuthentication(w http.ResponseWriter, r *http.Request) {
	var user model.UserAuthenticateRequest
	if r.Body == nil {
		util.WriteJSON(w, 400, model.Response{
			Code:   400,
			Status: "BAD_REQUEST",
			Error:  "Missing body request",
		})
		return
	}
	if err := util.ParseJSON(r, &user); err != nil {
		util.WriteJSON(w, 400, model.Response{
			Code:   400,
			Status: "BAD_REQUEST",
			Error:  "err.Error()",
		})
		return
	}
	if err := h.validate.Struct(user); err != nil {
		util.WriteJSON(w, 422, model.Response{
			Code:   422,
			Status: "UNPROCESSABLE_ENTITY",
			Error:  err.Error(),
		})
		return
	}
	h.store.Authentication(w, r.Context(), user)
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var user model.UserRegisterRequest
	if r.Body == nil {
		util.WriteJSON(w, 400, model.Response{
			Code:   400,
			Status: "BAD_REQUEST",
			Error:  "Missing body request",
		})
		return
	}
	if err := util.ParseJSON(r, &user); err != nil {
		util.WriteJSON(w, 400, model.Response{
			Code:   400,
			Status: "BAD_REQUEST",
			Error:  err.Error(),
		})
		return
	}
	if err := h.validate.Struct(user); err != nil {
		util.WriteJSON(w, 422, model.Response{
			Code:   422,
			Status: "UNPROCESSABLE_ENTITY",
			Error:  err.Error(),
		})
		return
	}
	h.store.Register(w, r.Context(), user)
}

func (h *Handler) Fetch(w http.ResponseWriter, r *http.Request) {
	h.store.Fetch(w, r.Context())
}

func (h *Handler) Modification(w http.ResponseWriter, r *http.Request) {
	var user model.UserModifyRequest
	if r.Body == nil {
		util.WriteJSON(w, 400, model.Response{
			Code:   400,
			Status: "BAD_REQUEST",
			Error:  "Missing body request",
		})
		return
	}
	if err := util.ParseJSON(r, &user); err != nil {
		util.WriteJSON(w, 400, model.Response{
			Code:   400,
			Status: "BAD_REQUEST",
			Error:  err.Error(),
		})
		return
	}
	if err := h.validate.Struct(user); err != nil {
		util.WriteJSON(w, 422, model.Response{
			Code:   422,
			Status: "UNPROCESSABLE_ENTITY",
			Error:  err.Error(),
		})
		return
	}
	h.store.Modification(w, r.Context(), user)
}
