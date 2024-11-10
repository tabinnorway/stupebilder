package albums

import (
	"log"

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

func (s Store) Create(data models.Album) (*models.Album, error) {
	newId := utils.CreateShortUUID()
	sql := `insert into albums(
			id,
			created_at,
			album_folder,
			title,
			datestring
		) values (
		 	$1, current_timestamp, $2, $3, $4
		)`
	_, err := s.db.Exec(sql, newId, data.AlbumFolder, data.Title, data.Datestring)
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
	// user, err := s.GetByID(id)
	// if err != nil {
	// 	return nil, err
	// }
	// if user == nil {
	// 	return nil, fmt.Errorf("not found")
	// }
	// sql := "DELETE FROM users where id = $1"
	// _, err = s.db.Exec(sql, id)
	// if err != nil {
	// 	return nil, err
	// }
	// return s.GetByID((id))
	return nil, nil
}

func (s *Store) Update(id string, user *models.User) (*models.User, error) {
	// sql := `update users set updated_at = current_timestamp,
	// 			email = $1,
	// 			passwd = $2,
	// 			username = $3,
	// 			first_name = $4,
	// 			last_name = $5,
	// 			primary_phone = $6,
	// 			primary_club_id = $7
	// 		where id = $8`
	// _, err := s.db.Exec(sql, user.Email, user.Password, user.Username, user.FirstName, user.LastName, user.PrimaryPhone, user.PrimaryClubId, id)
	// if err != nil {
	// 	return nil, err
	// }
	// return s.GetByID((id))
	return nil, nil
}
