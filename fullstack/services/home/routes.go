package home

import (
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/tabinnorway/stupebilder/utils"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Get("/", h.getHomePage)
	r.Get("/style/main.css", h.getCss)
}

func (h *Handler) getHomePage(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "index.html", utils.PageData{Title: "Please enter key"})
}

func (h *Handler) getCss(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("style", "main.css")
	w.Header().Set("Content-Type", "text/css")
	http.ServeFile(w, r, filePath)
}
