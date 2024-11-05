package albums

import (
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
	r.Get("/", h.getAlbums)
}

func (h *Handler) getAlbums(w http.ResponseWriter, r *http.Request) {
	templateName := "albums.html"
	foldersPath := h.imgRoot

	var files []utils.FileInfo
	entries, err := os.ReadDir(foldersPath)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, nil)
	}
	for _, entry := range entries {
		if !strings.HasPrefix(entry.Name(), ".") && entry.IsDir() {
			albumInfo := utils.GetAlbumInfo(filepath.Join(foldersPath, entry.Name()))
			files = append(files, utils.FileInfo{
				Title:    albumInfo.Title,
				Date:     albumInfo.Date,
				Location: albumInfo.Location,
				Name:     entry.Name(),
				IsDir:    entry.IsDir(),
			})
		}
	}

	data := utils.PageData{
		Title:  "Albums",
		Album:  "",
		Folder: "",
		Files:  files,
	}
	utils.RenderTemplate(w, templateName, data)
}
