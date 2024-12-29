package handlers

import (
	"encoding/json"
	"net/http"
	"rent-server/internal/api/middleware"
	"rent-server/internal/services"
	"strconv"
)

type UserHandler struct {
	Service *services.UserService
}

// costructor
func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	value := r.Context().Value(middleware.GetKey())
	userIDStr, ok := value.(string)
    if !ok {
        http.Error(w, "invalid user ID", http.StatusBadRequest)
        return
    }

	userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        http.Error(w, "invalid user ID format", http.StatusBadRequest)
        return
    }
	
	user, err := h.Service.GetUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
