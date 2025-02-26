package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gagoto-dev/aevoleyback/internal/db/repositories"
)

type UserHandler struct {
	userRepository *repositories.UserRepository
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{
		repositories.NewUserRepository(db),
	}
}

func (b *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

}

func (b *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"hello\": \"world\"}"))
}
