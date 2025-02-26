package routes

import (
	"net/http"

	"github.com/gagoto-dev/aevoleyback/internal/api/handlers"
)

func BlogMux(handler *handlers.BlogHandler) http.Handler {
	blogMux := http.NewServeMux()

	blogMux.HandleFunc("POST /create", handler.CreateBlogEntry)
	blogMux.HandleFunc("GET /", handler.GetBlogEntries)

	return http.StripPrefix("/api/v1/blog", blogMux)
}
