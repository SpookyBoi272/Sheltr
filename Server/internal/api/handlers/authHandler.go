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

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var data models.User
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	user, token, err := h.Service.LoginUser(&data)
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

func (h *AuthHandler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var data models.User
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}
	
	user, token, err := h.Service.SignupUser(&data)
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