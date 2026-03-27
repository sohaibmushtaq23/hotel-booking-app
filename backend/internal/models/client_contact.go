package models

import (
	"database/sql"
	"time"
)

type ClientContact struct {
	ID          int            `json:"id"`
	IDClient    int            `json:"idClient"`
	FirstName   string         `json:"firstName"`
	LastName    string         `json:"lastName"`
	Designation string         `json:"designation"`
	Email       string         `json:"email"`
	Mobile      string         `json:"mobile"`
	Gender      string         `json:"gender"`
	BirthDate   string         `json:"birthDate"`
	IsPrimary   bool           `json:"isPrimary"`
	Notes       sql.NullString `json:"notes"`
	CreatedAt   time.Time      `json:"createdAt"`
}
