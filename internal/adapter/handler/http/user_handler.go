package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/g-villarinho/hexagonal-demo/internal/adapter/handler/http/dto"
	"github.com/g-villarinho/hexagonal-demo/internal/core/domain"
	"github.com/g-villarinho/hexagonal-demo/internal/core/port"
)

type UserHandler struct {
	userService port.UserService
}

func NewUserHandler(userService port.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "corpo da requisição inválido", http.StatusBadRequest)
		return
	}

	user, err := h.userService.Create(r.Context(), req.Name, req.Email, req.Password)
	if err != nil {
		if errors.Is(err, domain.ErrEmailAlreadyExists) {
			http.Error(w, "Não foi possível criar um usuário com esse e-mail.", http.StatusConflict)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := dto.ToUserResponse(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "corpo da requisição inválido", http.StatusBadRequest)
		return
	}

	token, err := h.userService.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidCredentials) {
			http.Error(w, "email ou senha incorretos", http.StatusUnauthorized)
			return
		}
		http.Error(w, "ocorreu um erro interno", http.StatusInternalServerError)
		return
	}

	response := dto.LoginResponse{Token: token}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
