package models

import (
	"database/sql"
	"time"
)

type Country struct {
	Id                 int            `db:"id" json:"id"`
	CreatedAt          time.Time      `db:"created_at" json:"createdAt"`
	UpdatedAt          sql.NullTime   `db:"updated_at" json:"updatedAt"`
	Name               string         `db:"country_name" json:"name"`
	CountryCodeA2      sql.NullString `db:"country_code_a2" json:"countryCodeA2"`
	CountryCodeA3      sql.NullString `db:"country_code_a3" json:"countryCodeA3"`
	CountryCodeNumeric sql.NullString `db:"country_code_numeric" json:"countryCodeNumeric"`
	PhonePrefix        sql.NullString `db:"country_phone_prefix" json:"phonePrefix"`
	TLD                sql.NullString `db:"tld" json:"tld"`
}
