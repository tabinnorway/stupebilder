package models

import (
	"database/sql"
	"time"
)

type Folder struct {
	Id        string       `db:"id" json:"id"`
	CreatedAt time.Time    `db:"created_at" json:"createdAt"`
	CreatedBy string       `db:"created_by" json:"createdBy"`
	UpdatedAt sql.NullTime `db:"updated_at" json:"updatedAt"`
	UpdatedBy *string      `db:"updated_by" json:"updatedBy"`
	DeletedAt sql.NullTime `db:"deleted_at" json:"DeletedAt"`
	DeletedBy *string      `db:"deleted_by" json:"DeletedBy"`
	Title     string       `db:"title" json:"title"`
	Date      string       `db:"date" json:"date"`
	NumImages int          `db:"num_images" json:"numImages"`
}
