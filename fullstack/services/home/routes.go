package home

import (
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/tabinnorway/stupebilder/views"
)

type Handler struct {
	db *sqlx.DB
}

func NewHandler(db *sqlx.DB) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Get("/", h.getHomePage)
	r.Get("/style/main.css", h.getCss)
}

func (h *Handler) getHomePage(w http.ResponseWriter, r *http.Request) {
	views.Home().Render(r.Context(), w)
}

func (h *Handler) getCss(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("style", "main.css")
	w.Header().Set("Content-Type", "text/css")
	http.ServeFile(w, r, filePath)
}
