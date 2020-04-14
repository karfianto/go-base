// Package models contains application specific entities.
package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	//"github.com/pkg/errors"
)

// Report holds specific application settings linked to an Account.
type Report struct {
	// Reports represents public.reports
	ID        int       `json:"id"`
	AccountID int       // account_id
	UpdatedAt time.Time // date
	Complaint string    // complaint
}

// Validate validates Profile struct and returns validation errors.
func (p *Report) Validate() error {

	return validation.ValidateStruct(p,
		validation.Field(&p.AccountID, validation.Required),
	)
}
