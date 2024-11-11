package models

import (
	"database/sql"
	"time"
)

type Album struct {
	Id         string         `db:"id" json:"id"`
	CreatedAt  time.Time      `db:"created_at" json:"createdAt"`
	CreatedBy  string         `db:"created_by" json:"createdBy"`
	UpdatedAt  sql.NullTime   `db:"updated_at" json:"updatedAt"`
	UpdatedBy  *string        `db:"updated_by" json:"updatedBy"`
	DeletedAt  sql.NullTime   `db:"deleted_at" json:"DeletedAt"`
	DeletedBy  *string        `db:"deleted_by" json:"DeletedBy"`
	AlbumPath  string         `db:"album_path" json:"albumFolder"`
	Title      string         `db:"title" json:"title"`
	Datestring sql.NullString `db:"datestring" json:"datestring"`
	OwnerId    sql.NullInt32  `db:"owner_id" json:"ownerId"`
}

type AlbumCreateDTO struct {
	AlbumPath  string  `db:"album_path" json:"albumFolder"`
	Title      string  `db:"title" json:"title"`
	Datestring *string `db:"datestring" json:"datestring"`
}

func (m *AlbumCreateDTO) ToModel() *Album {
	album := Album{
		Id:         "",
		CreatedAt:  time.Time{},
		UpdatedAt:  sql.NullTime{Valid: false},
		AlbumPath:  m.AlbumPath,
		Title:      m.Title,
		Datestring: sql.NullString{Valid: false},
	}
	return &album
}
