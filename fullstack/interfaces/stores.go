package interfaces

import "github.com/tabinnorway/stupebilder/models"

type UserStore interface {
	GetAll() ([]models.User, error)
	GetByID(id string) (*models.User, error)
	Create(models.User) (*models.User, error)
	Delete(id string) (*models.User, error)
	Update(id string, user *models.User) (*models.User, error)
	GetBuEmail(email string) (*models.User, error)
}

type AlbumStore interface {
	GetAll() ([]models.Album, error)
	GetByID(string) (*models.Album, error)
	Create(models.Album) (*models.Album, error)
	GetThumb(id string) string
}

type FolderStore interface {
	GetAll() ([]models.Folder, error)
	GetByID(string, string) (*models.Folder, error)
	GetThumb(string, string) string
}
