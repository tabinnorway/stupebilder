package models

import (
	"database/sql"
	"time"
)

type Club struct {
	Id               string         `db:"id" json:"id"`
	CreatedAt        time.Time      `db:"created_at" json:"createdAt"`
	UpdatedAt        sql.NullTime   `db:"updated_at" json:"updatedAt"`
	Email            string         `db:"email" json:"email"`
	ClubName         string         `db:"club_name" json:"clubName"`
	ShortName        string         `db:"short_name" json:"shortName"`
	PhoneNumber      sql.NullString `db:"phone_number" json:"phoneNumber"`
	StreetAddress    sql.NullString `db:"street_address" json:"streetAddress"`
	PostalCode       sql.NullString `db:"postal_code" json:"postalCode"`
	City             sql.NullString `db:"city" json:"city"`
	CountryId        sql.NullInt32  `db:"country_id" json:"countryId"`
	PrimaryContactId sql.NullInt32  `db:"primary_contact_id" json:"primaryContactId"`
}
