package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gagoto-dev/aevoleyback/internal/api/handlers"
	"github.com/gagoto-dev/aevoleyback/internal/api/routes"
	"github.com/rs/cors"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	mux := http.NewServeMux()

	// Handlers
	blogHandler := handlers.NewBlogHandler(s.db)
	userHandler := handlers.NewUserHandler(s.db)

	mux.Handle("/api/v1/blog/", routes.BlogMux(blogHandler))
	mux.Handle("/api/v1/user/", routes.UsersMux(userHandler))

	// TODO: Middleware

	corsHandler := cors.AllowAll().Handler(mux)

	log.Printf("API Running on PORT %v\n", s.addr)

	err := http.ListenAndServe(s.addr, corsHandler)
	return err
}
