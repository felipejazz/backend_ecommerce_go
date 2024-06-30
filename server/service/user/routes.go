package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/felipejazz/ecommerce_go/config"
	"github.com/felipejazz/ecommerce_go/service/auth"
	"github.com/felipejazz/ecommerce_go/types"
	"github.com/felipejazz/ecommerce_go/utils"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/hello", h.handleHello).Methods("GET")
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")

}

func (h *Handler) handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// GET THE JSON PAYLOAD
	var payload types.LoginUserPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	if err := utils.Validate.Struct(payload); err != nil {
		error := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", error))
		return
	}

	u, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		log.Printf("%s", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return
	}

	if !auth.ComparePasswords(u.Password, []byte(payload.Password)) {
		log.Printf("Password mismatch: stored hash: %s, provided password: %s\n", u.Password, payload.Password)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, u.ID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	utils.WriteJson(w, http.StatusOK, map[string]string{"token": token})

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// GET THE JSON PAYLOAD
	var payload types.RegisterUserPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	if err := utils.Validate.Struct(payload); err != nil {
		error := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", error))
		return
	}

	// CHECK IF USER EXISTS
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, nil)

}
