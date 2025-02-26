package routes

import (
	"net/http"

	"github.com/gagoto-dev/aevoleyback/internal/api/handlers"
)

func UsersMux(handler *handlers.UserHandler) http.Handler {
	usersMux := http.NewServeMux()

	usersMux.HandleFunc("POST /create", handler.CreateUser)
	usersMux.HandleFunc("GET /", handler.GetUser)

	return http.StripPrefix("/api/v1/user", usersMux)
}
