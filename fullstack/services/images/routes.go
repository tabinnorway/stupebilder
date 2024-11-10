package images

import (
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/tabinnorway/stupebilder/interfaces"
	"github.com/tabinnorway/stupebilder/utils"
)

type Handler struct {
	albumStore  interfaces.AlbumStore
	folderStore interfaces.FolderStore
}

func NewHandler(as interfaces.AlbumStore, fs interfaces.FolderStore) *Handler {
	return &Handler{albumStore: as, folderStore: fs}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Get("/{albumId}/{folderId}/{imageFile}", h.getImage)
}

func (h *Handler) getImage(w http.ResponseWriter, r *http.Request) {
	albumId := chi.URLParam(r, "albumId")
	folderId := chi.URLParam(r, "folderId")
	imageFile := chi.URLParam(r, "imageFile")

	album, err := h.albumStore.GetByID(albumId)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, nil)
		return
	}
	path := filepath.Join(album.AlbumFolder, "images", folderId, imageFile)

	w.Header().Set("Content-Type", "image/jpeg")
	http.ServeFile(w, r, path)
}
