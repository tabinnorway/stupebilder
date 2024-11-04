package albums

import (
	"net/http"
	"os"
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
		files = append(files, utils.FileInfo{
			Name:  entry.Name(),
			IsDir: entry.IsDir(),
		})
	}

	filterFunc := func(s utils.FileInfo) bool {
		return s.IsDir && !strings.HasPrefix(s.Name, ".")
	}
	files = utils.Filter(files, filterFunc)
	data := utils.PageData{
		Title:  "Albums",
		Album:  "",
		Folder: "",
		Files:  files,
	}
	utils.RenderTemplate(w, templateName, data)
}
