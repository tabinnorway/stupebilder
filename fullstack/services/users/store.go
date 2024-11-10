package users

import (
	"fmt"
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

func (s Store) GetAll() ([]models.User, error) {
	var users []models.User
	query := `select u.*, c.club_name as "primary_club.club_name"
    			from users u
    			left outer join clubs c
				on u.primary_club_id = c.id`
	err := s.db.Select(&users, query)
	if err != nil {
		log.Printf("error getting users: %s", err.Error())
		return nil, err
	}
	return users, nil
}

func (s *Store) GetByID(id string) (*models.User, error) {
	user := models.User{}
	if err := s.db.Get(&user, "SELECT * FROM users WHERE id = $1", id); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *Store) GetBuEmail(email string) (*models.User, error) {
	user := models.User{}
	if err := s.db.Get(&user, "SELECT * FROM users WHERE lower(email) = lower($1)", email); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s Store) Create(user models.User) (*models.User, error) {
	newId := utils.CreateShortUUID()
	sql := `insert into users(
			id, created_at, email, passwd, username, first_name, last_name, primary_phone
		) values (
		 	$1 current_timestamp, $2, $3, $4, $5, $6, $7, $6
		)`
	_, err := s.db.Exec(sql, newId, user.Email, user.Password, user.Username, user.FirstName, user.LastName, user.PrimaryPhone)
	if err != nil {
		return nil, err
	}

	newUser, err := s.GetByID(newId)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (s *Store) Delete(id string) (*models.User, error) {
	user, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("not found")
	}
	sql := "DELETE FROM users where id = $1"
	_, err = s.db.Exec(sql, id)
	if err != nil {
		return nil, err
	}
	return s.GetByID((id))
}

func (s *Store) Update(id string, user *models.User) (*models.User, error) {
	sql := `
		update users set updated_at = current_timestamp,
						 email = $1,
						 passwd = $2,
						 username = $3,
						 first_name = $4,
						 last_name = $5,
						 primary_phone = $6,
						 primary_club_id = $7
					 where id = $8
		`
	_, err := s.db.Exec(sql, user.Email, user.Password, user.Username, user.FirstName, user.LastName, user.PrimaryPhone, user.PrimaryClubId, id)
	if err != nil {
		return nil, err
	}
	return s.GetByID((id))
}
