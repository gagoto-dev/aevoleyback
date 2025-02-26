package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gagoto-dev/aevoleyback/internal/db/repositories"
)

type BlogHandler struct {
	blogRepository *repositories.BlogEntryRepository
}

func NewBlogHandler(db *sql.DB) *BlogHandler {
	return &BlogHandler{
		repositories.NewBlogRepository(db),
	}
}

func (b *BlogHandler) CreateBlogEntry(w http.ResponseWriter, r *http.Request) {

}

func (b *BlogHandler) GetBlogEntries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"hello\": \"world\"}"))
}
