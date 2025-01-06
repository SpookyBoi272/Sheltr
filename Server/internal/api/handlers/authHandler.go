package handlers

import (
	"encoding/json"
	"net/http"
	"rent-server/internal/models"
	"rent-server/internal/services"
)

type AuthHandler struct{
	Service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler{
	return &AuthHandler{Service: service}
}

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid request data", http.StatusBadRequest)
        return
    }
	
	user := &models.User{
		Email: req.Email,
		Password: req.Password,
	}
	
	user, token, err := h.Service.LoginUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
        Value:    token,
        HttpOnly: true,
        Secure:   true,
        Path:     "/",
    })
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

type SignupRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (h *AuthHandler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req SignupRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	user := &models.User{
		Email: req.Email,
		Password: req.Password,
	}
	
	
	user, token, err := h.Service.SignupUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	http.SetCookie(w, &http.Cookie{
        Name:     "token",
        Value:    token,
        HttpOnly: true,
        Secure:   true,
        Path:     "/",
    })

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}