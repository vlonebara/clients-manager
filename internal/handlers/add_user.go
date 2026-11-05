package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// db connection
type UserHandler struct {
	DB *sql.DB
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{
		DB: db,
	}
}

// user request data
type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

// user response data
type CreateUserResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Login string `json:"login"`
}

// handler get user data in request and writing user in db
func (h *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var request CreateUserRequest

	//creating json decoder for reading data
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	//decoding request data
	err := decoder.Decode(&request)
	if err != nil {
		log.Println("[AddUserHandler]: Decode request json error: ", err)
	}

	//trimspase request data
	request.Name = strings.TrimSpace(request.Name)
	request.Email = strings.TrimSpace(request.Email)
	request.Role = strings.TrimSpace(request.Role)
	request.Login = strings.TrimSpace(request.Login)

	//generate password hash
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("[AddUserHandler]: Generate password hash error: ", err)
	}

	//writing data in db
	result, err := h.DB.ExecContext(
		r.Context(),
		`
		INSERT INTO users (
		name,
		email,
		role,
		login,
		password_hash)
		VALUES (
		?, ?, ?, ?, ?)
		`,
		request.Name,
		request.Email,
		request.Role,
		request.Login,
		string(passwordHash),
	)
	if err != nil {
		log.Println("[AddUserHandler]: Writing data in db error: ", err)
	}

	//get added user id
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("[AddUserHandler]: Ger added user id error: ", err)
	}

	//handler response
	response := CreateUserResponse{
		ID:    id,
		Name:  request.Name,
		Email: request.Email,
		Role:  request.Role,
		Login: request.Login,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(response)
}
