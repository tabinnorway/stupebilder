package albums

import (
	"fmt"
	"log"
	"os"
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
	entries, err := os.ReadDir(album.AlbumPath)
	if err != nil {
		return
	}

	for _, entry := range entries {
		if entry.IsDir() && !strings.HasPrefix(entry.Name(), ".") {
			fmt.Printf("Doing folder: %s\n", entry.Name())
			time.Sleep(1 * time.Second)
		}
	}
}

func (s Store) Create(data models.Album) (*models.Album, error) {
	if !utils.DirExists(data.AlbumPath) {
		return nil, fmt.Errorf("folder path does not exist")
	}
	go importFolder(data)
	fmt.Println("Started myFunction as a Goroutine")
	return nil, fmt.Errorf("not yet ready")

	// sql := `insert into albums(
	// 		id,
	// 		created_at,
	// 		album_path,
	// 		title,
	// 		datestring
	// 	) values (
	// 	 	$1, current_timestamp, $2, $3, $4
	// 	)`
	// newId := utils.CreateShortUUID()
	// _, err := s.db.Exec(sql, newId, data.AlbumPath, data.Title, data.Datestring)
	// if err != nil {
	// 	return nil, err
	// }

	// created, err := s.GetByID(newId)
	// if err != nil {
	// 	return nil, err
	// }
	// return created, nil
}

func (s *Store) Delete(id string) (*models.User, error) {
	return nil, nil
}

func (s *Store) Update(id string, user *models.User) (*models.User, error) {
	return nil, nil
}
