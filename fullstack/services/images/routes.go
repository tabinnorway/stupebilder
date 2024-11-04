package images

import (
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	imgRoot string
}

func NewHandler(imgRoot string) *Handler {
	return &Handler{imgRoot: imgRoot}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Get("/{albumId}/{folderId}/{image}", h.getImage)
}

func (h *Handler) getImage(w http.ResponseWriter, r *http.Request) {
	albumId := chi.URLParam(r, "albumId")
	folderId := chi.URLParam(r, "folderId")
	image := chi.URLParam(r, "image")
	thumbPath := filepath.Join(h.imgRoot, albumId, "images", folderId, image)

	w.Header().Set("Content-Type", "image/jpeg")
	http.ServeFile(w, r, thumbPath)
}
