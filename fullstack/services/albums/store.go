package albums

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/tabinnorway/stupebilder/models"
	"github.com/tabinnorway/stupebilder/utils"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{db: db}
}

func (s Store) GetAll() ([]models.Album, error) {
	var data []models.Album
	query := `select * from albums`
	err := s.db.Select(&data, query)
	if err != nil {
		log.Printf("error getting users: %s", err.Error())
		return nil, err
	}
	return data, nil
}

func (s *Store) GetByID(id string) (*models.Album, error) {
	data := models.Album{}
	if err := s.db.Get(&data, "select * from albums where id = $1", id); err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *Store) GetThumb(id string) string {
	return ""
}

func importFolder(album models.Album) {
	if !utils.DirExists(album.AlbumPath) {
		return
	}
	// Create an Images directory and a Thumbnails directory
	imagesPath := filepath.Join(album.AlbumPath, "Images")
	if !utils.DirExists(imagesPath) {
		if err := os.Mkdir(imagesPath, 0755); err != nil {
			log.Fatalf("Could not create images directory: %+v", err)
		}
	}
	thumbsPath := filepath.Join(album.AlbumPath, "Thumbs")
	if !utils.DirExists(thumbsPath) {
		if err := os.Mkdir(thumbsPath, 0755); err != nil {
			log.Fatalf("Could not create thumbs directory: %+v", err)
		}
	}

	entries, err := os.ReadDir(album.AlbumPath)
	if err != nil {
		return
	}
	for _, entry := range entries {
		if !strings.HasPrefix(entry.Name(), ".") && entry.Name() != "Images" && entry.Name() != "Thumbs" {
			if entry.IsDir() {
				fmt.Printf("Doing folder: %s\n", entry.Name())
				fromPath := filepath.Join(album.AlbumPath, entry.Name())
				toPath := filepath.Join(imagesPath, entry.Name())
				time.Sleep(1 * time.Second)
				if err = os.Rename(fromPath, toPath); err != nil {
					log.Fatalf("Could move folder to images : %+v", err)
				}
			}
		}
	}
	imgFolders, err := os.ReadDir(imagesPath)
	if err != nil {
		return
	}

	for _, imgFolder := range imgFolders {
		if !strings.HasPrefix(imgFolder.Name(), ".") && imgFolder.Name() != "Images" && imgFolder.Name() != "Thumbs" {
			theImagePath := filepath.Join(imagesPath, imgFolder.Name())
			theThumbImagePath := filepath.Join(thumbsPath, imgFolder.Name())
			if !utils.DirExists(theThumbImagePath) {
				if err := os.Mkdir(theThumbImagePath, 0755); err != nil {
					log.Fatalf("Could not create images directory: %+v", err)
				}
			}
			convertFolder(theImagePath, theThumbImagePath, 200)
		}
	}
}

func convertFolder(from string, to string, size int) {
	fmt.Printf("converting folder: %s\n", from)
	fmt.Printf("               to: %s\n", to)
	entries, err := os.ReadDir(from)
	if err != nil {
		log.Fatalf("Got an error reading dir: %+v", err)
	}
	for _, entry := range entries {
		if !strings.HasPrefix(entry.Name(), ".") && entry.Name() != "Images" && entry.Name() != "Thumbs" {
			fromPath := filepath.Join(from, entry.Name())
			toPath := filepath.Join(to, entry.Name())
			if entry.IsDir() {
				convertFolder(fromPath, toPath, size)
			} else {
				utils.ResizeImage(fromPath, toPath, size)
			}
		}
	}
}

func (s Store) Create(data models.Album) (*models.Album, error) {
	if !utils.DirExists(data.AlbumPath) {
		return nil, fmt.Errorf("folder path does not exist")
	}
	go importFolder(data)
	fmt.Println("Started myFunction as a Goroutine")

	sql := `insert into albums(
			id,
			created_at,
			created_by,
			album_path,
			title,
			datestring
		) values (
		 	$1, current_timestamp, (select id from users where email = 'terje@bergesen.info'), $2, $3, $4
		)`
	newId := utils.CreateShortUUID()
	_, err := s.db.Exec(sql, newId, data.AlbumPath, data.Title, data.Datestring)
	if err != nil {
		return nil, err
	}

	created, err := s.GetByID(newId)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (s *Store) Delete(id string) (*models.User, error) {
	return nil, nil
}

func (s *Store) Update(id string, user *models.User) (*models.User, error) {
	return nil, nil
}
