package user

import (
	"net/http"

	"github.com/felipejazz/ecommerce_go/cmd/types"
	"github.com/felipejazz/ecommerce_go/cmd/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

	// GET THE JSON PAYLOAD
	var payload types.RegisterUserPayload
	if err := utils.ParseJson(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// CHECK IF USER EXISTS

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

}
