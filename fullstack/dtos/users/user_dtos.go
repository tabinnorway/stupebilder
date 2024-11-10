package users

import (
	"database/sql"
	"time"

	"github.com/tabinnorway/stupebilder/models"
	"github.com/tabinnorway/stupebilder/utils"
)

type UserCreateDTO struct {
	Email        string  `json:"email"`
	Username     *string `json:"userName"`
	FisrstName   string  `json:"fisrstName"`
	LastName     string  `json:"lastName"`
	PrimaryPhone *string `json:"primaryPhone"`
}

func (dto *UserCreateDTO) ToModel() *models.User {
	model := models.User{
		Id:            "tempId",
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
