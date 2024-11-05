package images

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	imgRoot string
}

func NewHandler(imgRoot string) *Handler {
	return &Handler{imgRoot: imgRoot}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Get("/users", h.getUsers)
}

func (h *Handler) getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/jpeg")
	http.ServeFile(w, r, "")
}
