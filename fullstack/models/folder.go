package models

import (
	"database/sql"
	"time"
)

type Folder struct {
	Id        string       `db:"id" json:"id"`
	CreatedAt time.Time    `db:"created_at" json:"createdAt"`
	UpdatedAt sql.NullTime `db:"updated_at" json:"updatedAt"`
	Title     string       `db:"title" json:"title"`
	Date      string       `db:"date" json:"date"`
	NumImages int          `db:"num_images" json:"numImages"`
}
