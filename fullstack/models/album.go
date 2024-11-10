package models

import (
	"database/sql"
	"time"
)

type Album struct {
	Id          string         `db:"id" json:"id"`
	CreatedAt   time.Time      `db:"created_at" json:"createdAt"`
	UpdatedAt   sql.NullTime   `db:"updated_at" json:"updatedAt"`
	AlbumFolder string         `db:"album_folder" json:"albumFolder"`
	Title       string         `db:"title" json:"title"`
	Datestring  sql.NullString `db:"datestring" json:"datestring"`
	OwnerId     sql.NullInt32  `db:"owner_id" json:"ownerId"`
}

type AlbumCreateDTO struct {
	AlbumFolder string  `db:"album_folder" json:"albumFolder"`
	Title       string  `db:"title" json:"title"`
	Datestring  *string `db:"datestring" json:"datestring"`
}

func (m *AlbumCreateDTO) ToModel() *Album {
	album := Album{
		Id:          "",
		CreatedAt:   time.Time{},
		UpdatedAt:   sql.NullTime{Valid: false},
		AlbumFolder: m.AlbumFolder,
		Title:       m.Title,
		Datestring:  sql.NullString{Valid: false},
	}
	return &album
}
