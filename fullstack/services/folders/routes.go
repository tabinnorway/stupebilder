package folders

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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
	r.Get("/{albumId}", h.getFolders)
	r.Get("/{albumId}/{folderId}", h.getfolderByName)
	r.Get("/{albumId}/{folderId}/download", h.downloadZip)
}

func (h *Handler) downloadZip(w http.ResponseWriter, r *http.Request) {
	albumId := chi.URLParam(r, "albumId")
	folderId := chi.URLParam(r, "folderId")
	foldersPath := filepath.Join(h.imgRoot, albumId, "images", folderId)
	zipFile := filepath.Join(foldersPath, "images.zip")
	if _, err := os.Stat(zipFile); err == nil {
		contentDisposition := fmt.Sprintf("attachment; filename=%s-images.zip", folderId)
		w.Header().Set("Content-Disposition", contentDisposition)
		http.ServeFile(w, r, zipFile)
		return
	}
	utils.RenderTemplate(w, "not_yet_ready.html", utils.PageData{})
}

func (h *Handler) getFolders(w http.ResponseWriter, r *http.Request) {
	albumId := chi.URLParam(r, "albumId")
	foldersPath := filepath.Join(h.imgRoot, albumId, "images")

	var directories []utils.FileInfo
	entries, err := os.ReadDir(foldersPath)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, nil)
	}
	for _, entry := range entries {
		if entry.IsDir() && !strings.HasPrefix(entry.Name(), ".") {
			directories = append(directories, utils.FileInfo{
				Name:  entry.Name(),
				IsDir: entry.IsDir(),
			})
		}
	}

	data := utils.PageData{
		Title:  albumId,
		Album:  albumId,
		Folder: "",
		Files:  directories,
	}
	utils.RenderTemplate(w, "folders.html", data)
}

func (h *Handler) getfolderByName(w http.ResponseWriter, r *http.Request) {
	templateName := "folder.html"
	albumId := chi.URLParam(r, "albumId")
	folderId := chi.URLParam(r, "folderId")
	foldersPath := filepath.Join(h.imgRoot, albumId, "images", folderId)
	entries, err := os.ReadDir(foldersPath)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, nil)
	}

	var files []utils.FileInfo
	for _, entry := range entries {
		files = append(files, utils.FileInfo{
			Name:  entry.Name(),
			IsDir: entry.IsDir(),
		})
	}
	filterFunc := func(s utils.FileInfo) bool {
		return !s.IsDir && !strings.HasPrefix(s.Name, ".") && !strings.HasSuffix(s.Name, ".zip")
	}
	files = utils.Filter(files, filterFunc)
	data := utils.PageData{
		Title:  folderId,
		Album:  albumId,
		Folder: folderId,
		Files:  files,
	}
	utils.RenderTemplate(w, templateName, data)
}
