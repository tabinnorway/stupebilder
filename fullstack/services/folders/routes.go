package folders

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/tabinnorway/stupebilder/interfaces"
	"github.com/tabinnorway/stupebilder/models"
	"github.com/tabinnorway/stupebilder/utils"
	"github.com/tabinnorway/stupebilder/views"
)

type Handler struct {
	albumStore  interfaces.AlbumStore
	folderStore interfaces.FolderStore
}

func NewHandler(as interfaces.AlbumStore, fs interfaces.FolderStore) *Handler {
	return &Handler{albumStore: as, folderStore: fs}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Get("/{albumId}/{folderId}", h.getFolderById)
	r.Get("/{albumId}/{folderId}/download", h.downloadZip)
}

func (h *Handler) downloadZip(w http.ResponseWriter, r *http.Request) {
	albumId := chi.URLParam(r, "albumId")
	folderId := chi.URLParam(r, "folderId")
	if len(albumId) <= 0 || len(folderId) <= 0 {
		utils.WriteError(w, http.StatusBadRequest, nil)
		return
	}
	album, err := h.albumStore.GetByID(albumId)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, nil)
		return
	}

	foldersPath := filepath.Join(album.AlbumPath, "images", folderId)
	zipFile := filepath.Join(foldersPath, "images.zip")
	if _, err := os.Stat(zipFile); err == nil {
		contentDisposition := fmt.Sprintf("attachment; filename=%s-images.zip", folderId)
		w.Header().Set("Content-Disposition", contentDisposition)
		http.ServeFile(w, r, zipFile)
		return
	}
	utils.WriteError(w, http.StatusNotFound, nil)
}

func (h *Handler) getFolderById(w http.ResponseWriter, r *http.Request) {
	albumId := chi.URLParam(r, "albumId")
	folderId := chi.URLParam(r, "folderId")

	if len(albumId) <= 0 || len(folderId) <= 0 {
		utils.WriteError(w, http.StatusBadRequest, nil)
		return
	}

	album, err := h.albumStore.GetByID(albumId)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, nil)
		return
	}
	foldersPath := filepath.Join(album.AlbumPath, "images", folderId)
	folder := models.Folder{Id: folderId, Title: folderId}
	fmt.Println(foldersPath)
	entries, err := os.ReadDir(foldersPath)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, nil)
		return
	}

	var images []models.Image
	for _, entry := range entries {
		if strings.HasSuffix(entry.Name(), "jpg") && !strings.HasPrefix(entry.Name(), ".") {
			images = append(images, models.Image{
				Id:        entry.Name(),
				CreatedAt: time.Now(),
				UpdatedAt: sql.NullTime{Valid: false},
				Title:     entry.Name(),
				Date:      "",
				FileName:  entry.Name(),
			})
		}
	}
	views.Folder(album, &folder, images).Render(r.Context(), w)
}
