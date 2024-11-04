package thumbs

import (
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/tabinnorway/stupebilder/utils"
)

type Handler struct {
	imgRoot string
}

func NewHandler(imgRoot string) *Handler {
	return &Handler{imgRoot: imgRoot}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Get("/{albumId}", h.getAlbumThumb)
	r.Get("/{albumId}/{folderId}", h.getFolderThumb)
	r.Get("/{albumId}/{folderId}/{image}", h.getThumb)
}

func (h *Handler) getFolderThumb(w http.ResponseWriter, r *http.Request) {
	albumId := chi.URLParam(r, "albumId")
	folderId := chi.URLParam(r, "folderId")

	albumPath := filepath.Join(h.imgRoot, albumId)
	file := utils.FindFolderThumb(albumPath, folderId)

	w.Header().Set("Content-Type", "image/jpeg")
	http.ServeFile(w, r, file)
}

func (h *Handler) getAlbumThumb(w http.ResponseWriter, r *http.Request) {
	albumId := chi.URLParam(r, "albumId")
	albumPath := filepath.Join(h.imgRoot, albumId)
	thumbPath := utils.FindAlbumThub(albumPath)

	w.Header().Set("Content-Type", "image/jpeg")
	http.ServeFile(w, r, thumbPath)
}

func (h *Handler) getThumb(w http.ResponseWriter, r *http.Request) {
	albumId := chi.URLParam(r, "albumId")
	folderId := chi.URLParam(r, "folderId")
	image := chi.URLParam(r, "image")
	thumbPath := filepath.Join(h.imgRoot, albumId, "thumbnails", folderId, image)

	w.Header().Set("Content-Type", "image/jpeg")
	http.ServeFile(w, r, thumbPath)
}
