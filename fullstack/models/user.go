package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id            string         `db:"id" json:"id"`
	CreatedAt     time.Time      `db:"created_at" json:"createdAt"`
	UpdatedAt     sql.NullTime   `db:"updated_at" json:"updatedAt"`
	Email         string         `db:"email" json:"email"`
	Username      sql.NullString `db:"username" json:"username"`
	Password      sql.NullString `db:"passwd" json:"password"`
	FirstName     string         `db:"first_name" json:"firstName"`
	LastName      string         `db:"last_name" json:"lastName"`
	PrimaryPhone  sql.NullString `db:"primary_phone" json:"primaryPhone"`
	Confirmed     bool           `db:"confirmed" json:"confirmed"`
	PrimaryClubId sql.NullString `db:"primary_club_id" json:"primaryClubId"`
	PrimaryClub   *Club          `db:"primary_club" json:"primaryClub"`
}
