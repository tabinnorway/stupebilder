package models

import (
	"database/sql"
	"time"

	"github.com/tabinnorway/stupebilder/utils"
)

type User struct {
	Id             string         `db:"id" json:"id"`
	CreatedAt      time.Time      `db:"created_at" json:"createdAt"`
	CreatedByEmail string         `db:"created_by_email" json:"createdByEmail"`
	UpdatedAt      sql.NullTime   `db:"updated_at" json:"updatedAt"`
	UpdatedBy      *string        `db:"updated_by" json:"updatedBy"`
	DeletedAt      sql.NullTime   `db:"deleted_at" json:"DeletedAt"`
	DeletedBy      *string        `db:"deleted_by" json:"DeletedBy"`
	Email          string         `db:"email" json:"email"`
	Username       sql.NullString `db:"username" json:"username"`
	Password       sql.NullString `db:"passwd" json:"password"`
	FirstName      string         `db:"first_name" json:"firstName"`
	LastName       string         `db:"last_name" json:"lastName"`
	PrimaryPhone   sql.NullString `db:"primary_phone" json:"primaryPhone"`
	Confirmed      bool           `db:"confirmed" json:"confirmed"`
	PrimaryClubId  sql.NullString `db:"primary_club_id" json:"primaryClubId"`
	PrimaryClub    *Club          `db:"primary_club" json:"primaryClub"`
}

type UserCreateDTO struct {
	Email        string  `json:"email"`
	Username     *string `json:"userName"`
	FisrstName   string  `json:"fisrstName"`
	LastName     string  `json:"lastName"`
	PrimaryPhone *string `json:"primaryPhone"`
}

func (dto *UserCreateDTO) ToModel() *User {
	model := User{
		Id:            "",
		CreatedAt:     time.Now(),
		UpdatedAt:     sql.NullTime{Valid: false},
		Email:         dto.Email,
		Username:      utils.NullString(dto.Username),
		Password:      utils.NullString(nil),
		FirstName:     dto.FisrstName,
		LastName:      dto.LastName,
		PrimaryPhone:  utils.NullString(dto.PrimaryPhone),
		Confirmed:     false,
		PrimaryClubId: utils.NullString(nil),
		PrimaryClub:   nil,
	}
	return &model
}
