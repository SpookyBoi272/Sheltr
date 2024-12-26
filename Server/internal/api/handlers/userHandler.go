package handlers

import (
	"net/http"
	"rent-server/internal/services"
)

type UserHandler struct {
	Service *services.UserService
}

//costructor
func NewUserHandler(service *services.UserService) *UserHandler{
	return &UserHandler{Service: service}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request){
	// vars := mux.Vars(r)

	// id ,err := strconv.Atoi(vars["id"])
	// if err != nil {
	// 	http.Error(w, "Invalid user ID", http.StatusBadRequest)
	// 	return
	// }

	// user, err := h.Service.GetUser(id)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusNotFound)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter , r *http.Request){
	// users, err := h.Service.GetAllUsers()
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusNotFound)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(users)
}