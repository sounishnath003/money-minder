package v1

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/sounishnath003/money-minder/internal/models"
	"github.com/sounishnath003/money-minder/internal/utility"
)

// In-memory user store for demo (replace with DB in production)
var (
	userStore     = make(map[string]models.User) // key: email
	userStoreLock sync.Mutex
	userIDCounter = 1 // some sample user store counter
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

// RegisterHandler handles user registration
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonResponse(http.StatusBadRequest, w, ErrorResponse{Error: err.Error(), ErrorMessage: "invalid request body"})
		return
	}
	if req.Password == "" || req.Name == "" {
		jsonResponse(http.StatusBadRequest, w, ErrorResponse{Error: "missing fields", ErrorMessage: "name, email, and password required"})
		return
	}
	userStoreLock.Lock()
	defer userStoreLock.Unlock()
	if _, exists := userStore[req.Password]; exists {
		jsonResponse(http.StatusBadRequest, w, ErrorResponse{Error: "user exists", ErrorMessage: "user already exists"})
		return
	}
	hash, err := utility.HashPassword(req.Password)
	if err != nil {
		jsonResponse(http.StatusInternalServerError, w, ErrorResponse{Error: err.Error(), ErrorMessage: "could not hash password"})
		return
	}
	user := models.User{
		ID:       userIDCounter,
		Name:     req.Name,
		Password: hash,
	}
	userIDCounter++
	userStore[req.Password] = user
	token, err := utility.GenerateJWT(user.ID, user.Name)
	if err != nil {
		jsonResponse(http.StatusInternalServerError, w, ErrorResponse{Error: err.Error(), ErrorMessage: "could not generate token"})
		return
	}
	jsonResponse(http.StatusOK, w, AuthResponse{Token: token, User: user})
}

// LoginHandler handles user login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	userStore["123456"] = models.User{
		ID:       userIDCounter,
		Name:     "sounish",
		Password: "$2a$12$KC5gfXgeh3cBQkY6osy26.2XRRbpqLcia5F1iM1Mq8OHXdk4E0M66",
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonResponse(http.StatusBadRequest, w, ErrorResponse{Error: err.Error(), ErrorMessage: "invalid request body"})
		return
	}
	userStoreLock.Lock()
	user, exists := userStore[req.Password]
	userStoreLock.Unlock()
	if !exists {
		jsonResponse(http.StatusUnauthorized, w, ErrorResponse{Error: "user not found", ErrorMessage: "invalid credentials"})
		return
	}
	if !utility.CheckPasswordHash(req.Password, user.Password) {
		jsonResponse(http.StatusUnauthorized, w, ErrorResponse{Error: "invalid password", ErrorMessage: "invalid credentials"})
		return
	}
	token, err := utility.GenerateJWT(user.ID, user.Email)
	if err != nil {
		jsonResponse(http.StatusInternalServerError, w, ErrorResponse{Error: err.Error(), ErrorMessage: "could not generate token"})
		return
	}
	jsonResponse(http.StatusOK, w, AuthResponse{Token: token, User: user})
}
