package albums

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/tabinnorway/stupebilder/interfaces"
	"github.com/tabinnorway/stupebilder/models"
	"github.com/tabinnorway/stupebilder/utils"
	"github.com/tabinnorway/stupebilder/views"
)

type Handler struct {
	store interfaces.AlbumStore
}

func NewHandler(s interfaces.AlbumStore) *Handler {
	return &Handler{store: s}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Get("/", h.getAlbums)
	r.Post("/", h.createAlbum)
	r.Get("/{id}/thumb", h.getAlbumThumb)
	r.Put("/{id}", h.updateAlbum)
	r.Get("/{id}/folders", h.getAlbumFolders)
	r.Get("/{id}/folders/{folderName}/thumb", h.getAlbumFolderThumb)
}

func (h *Handler) getAlbums(w http.ResponseWriter, r *http.Request) {
	albums, err := h.store.GetAll()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	views.Albums(&albums).Render(r.Context(), w)
}

func (h *Handler) createAlbum(w http.ResponseWriter, r *http.Request) {
	var data models.AlbumCreateDTO
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	newUser, err := h.store.Create(*data.ToModel())
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, newUser)

}

func (h *Handler) updateAlbum(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusNotFound, nil)
}

func (h *Handler) getAlbumThumb(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	album, err := h.store.GetByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, nil)
	}
	thumbPath := utils.FindAlbumThub(filepath.Join(album.AlbumPath, "Thumbs"))
	w.Header().Set("Content-Type", "image/jpeg")
	http.ServeFile(w, r, thumbPath)
}

func (h *Handler) getAlbumFolders(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if len(id) <= 0 {
		utils.WriteError(w, http.StatusBadRequest, nil)
		return
	}

	album, err := h.store.GetByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, nil)
	}
	foldersPath := filepath.Join(album.AlbumPath, "images")
	entries, err := os.ReadDir(foldersPath)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, nil)
		return
	}
	var folders []models.Folder
	for _, entry := range entries {
		if entry.IsDir() && !strings.HasPrefix(entry.Name(), ".") {
			folders = append(folders, models.Folder{
				Title:     entry.Name(),
				Date:      album.Datestring.String,
				NumImages: 0,
			})
		}
	}
	views.Folders(album, &folders).Render(r.Context(), w)
}

func (h *Handler) getAlbumFolderThumb(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	folderName := chi.URLParam(r, "folderName")
	if len(folderName) <= 0 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing folder pram"))
		return
	}

	album, err := h.store.GetByID(id)
	if err != nil || album == nil {
		utils.WriteError(w, http.StatusNotFound, nil)
		return
	}

	file := utils.FindFolderThumb(album.AlbumPath, folderName)
	w.Header().Set("Content-Type", "image/jpeg")
	http.ServeFile(w, r, file)
}
