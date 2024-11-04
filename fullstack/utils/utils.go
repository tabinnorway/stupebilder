package utils

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/google/uuid"
	resizer "github.com/nfnt/resize"
)

// var Validate = validator.New()

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	if err != nil {
		WriteJSON(w, status, map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(status)
}

func ParseJSON(r *http.Request, v any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(v)
}

func GetTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	tokenQuery := r.URL.Query().Get("token")

	if tokenAuth != "" {
		return tokenAuth
	}

	if tokenQuery != "" {
		return tokenQuery
	}

	return ""
}

func ParseID(idString string) int64 {
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return -1
	}
	return id
}

func DirExists(dirname string) bool {
	info, err := os.Stat(dirname)
	return err == nil && info.IsDir()
}

func EnsureDir(dirName string) error {
	// Check if the directory exists
	info, err := os.Stat(dirName)
	if err == nil || info.IsDir() {
		return nil
	}

	if os.IsNotExist(err) {
		// Directory does not exist, create it
		err := os.MkdirAll(dirName, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	} else if !info.IsDir() {
		return fmt.Errorf("path exists but is not a directory: %s", dirName)
	}
	return nil
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func CreateImage(fromDir string, toDir string, filename string, which string) error {
	fromPath := fmt.Sprintf("%s/%s", fromDir, filename)
	toPath := fmt.Sprintf("%s/%s", toDir, filename)

	file, err := os.Open(fromPath)
	if err != nil {
		return fmt.Errorf("failed to open input file: %w", err)
	}
	defer file.Close()

	imgConfig, _, err := image.DecodeConfig(file)
	if err != nil {
		return err
	}
	newLength := float64(200)
	if which == "preview" {
		newLength = 1024
	}

	factor := newLength / float64(imgConfig.Width)

	width := float64(imgConfig.Width) * factor
	height := uint(float64(imgConfig.Height) * factor)
	newWidth := uint(width)

	// Reset filepointer before we try to decode
	file.Seek(0, io.SeekStart)

	img, err := jpeg.Decode(file)
	if err != nil {
		return fmt.Errorf("failed to decode JPEG image: %w", err)
	}
	newImage := resizer.Resize(newWidth, uint(height), img, resizer.Lanczos3)
	outFile, err := os.Create(toPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outFile.Close()

	err = jpeg.Encode(outFile, newImage, nil)
	if err != nil {
		return fmt.Errorf("failed to encode resized image: %w", err)
	}
	return nil
}

func CreateZip(files []string, inFolder string, output string) (string, error) {
	// Create a zip file
	zipFile, err := os.Create(output)
	if err != nil {
		id := uuid.New()
		output = fmt.Sprintf("/tmp/%s.zip", id)
		zipFile, err = os.Create(output)
		if err != nil {
			return "", err
		}
	}
	defer zipFile.Close()

	// Create a new zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Add files to the zip
	for _, file := range files {
		file = fmt.Sprintf("%s/%s", inFolder, file)
		if err := addFileToZip(zipWriter, file); err != nil {
			return "", err
		}
	}
	return output, nil
}

// addFileToZip adds an individual file to the zip archive
func addFileToZip(zipWriter *zip.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get file info to use in zip header
	info, err := file.Stat()
	if err != nil {
		return err
	}

	// Create a zip header based on file info
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	header.Name = filepath.Base(filename)
	header.Method = zip.Deflate // Use Deflate compression

	// Create a writer for this file in the zip
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	// Copy file content to the zip
	_, err = io.Copy(writer, file)
	return err
}

func Filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

type FileInfo struct {
	Name  string
	IsDir bool
}

type PageData struct {
	Title  string
	Album  string
	Folder string
	Files  []FileInfo
}

func FindAlbumThub(albumPath string) string {
	return filepath.Join(albumPath, "thumb.jpg")
}

func FindFolderThumb(albumPath string, folderName string) string {
	path := filepath.Join(albumPath, "thumbnails", folderName)
	files, err := os.ReadDir(path)
	if err != nil {
		return " /mnt/familyshare/images/generic-thumb.jpg"
	}
	for _, elem := range files {
		if !elem.IsDir() && strings.HasSuffix(elem.Name(), "jpg") {
			return filepath.Join(albumPath, "thumbnails", folderName, elem.Name())
		}
	}
	return " /mnt/familyshare/images/generic-thumb.jpg"
}

// renderTemplate parses and executes templates with a common layout
func RenderTemplate(w http.ResponseWriter, tmpl string, data PageData) {
	templates, err := template.ParseFiles(
		"templates/layout.html",
		filepath.Join("templates", tmpl),
	)
	if err != nil {
		http.Error(w, "Error loading templates", http.StatusInternalServerError)
		log.Println("Template parsing error:", err)
		return
	}

	err = templates.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Println("Template execution error:", err)
	}
}
