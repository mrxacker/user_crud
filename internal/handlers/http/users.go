package httpHandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mrxacker/user_service/internal/dto"
	"github.com/mrxacker/user_service/internal/models"
)

type UserService interface {
	RegisterUser(name, email string) (models.User, error)
	GetUser(id int) (models.User, error)
	GetUsers() ([]models.User, error)
}

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", h.GetUsers)
	r.Post("/", h.CreateUser)
	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.GetUserByID)
		r.Put("/", h.GetUserByID)
		r.Delete("/", h.GetUserByID)
	})
	return r
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		jsonError(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	created, err := h.userService.RegisterUser(user.Name, user.Email)
	if err != nil {
		jsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse(w, created, http.StatusCreated)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.GetUsers()
	if err != nil {
		jsonError(w, "users not found", http.StatusNotFound)
		return
	}

	jsonResponse(w, users, http.StatusOK)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		jsonError(w, "invalid id", http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUser(id)
	if err != nil {
		jsonError(w, "user not found", http.StatusNotFound)
		return
	}

	jsonResponse(w, user, http.StatusOK)
}

func jsonResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return
	}
}

func jsonError(w http.ResponseWriter, msg string, status int) {
	jsonResponse(w, map[string]string{"error": msg}, status)
}
