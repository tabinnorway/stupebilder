package models

import (
	"database/sql"
	"time"
)

type Image struct {
	Id        string       `db:"id" json:"id"`
	CreatedAt time.Time    `db:"created_at" json:"createdAt"`
	UpdatedAt sql.NullTime `db:"updated_at" json:"updatedAt"`
	Title     string       `db:"title" json:"title"`
	Date      string       `db:"date" json:"date"`
	FileName  string       `db:"file_name" json:"fileName"`
}
