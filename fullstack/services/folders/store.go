package folders

import (
	"github.com/jmoiron/sqlx"
	"github.com/tabinnorway/stupebilder/models"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{db: db}
}

func (s Store) GetAll() ([]models.Folder, error) {
	return nil, nil
}

func (s Store) GetByID(albumId string, folderId string) (*models.Folder, error) {
	return nil, nil
}

func (s Store) GetThumb(albumId string, folderId string) string {
	return ""
}
