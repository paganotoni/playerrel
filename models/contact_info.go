package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type ContactInfo struct {
	ID          uuid.UUID `json:"id" db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	ParentName  string    `json:"parent_name" db:"parent_name"`
	PhoneNumber string    `json:"phone_number" db:"phone_number"`

	PlayerID uuid.UUID `json:"player_id" db:"p_id"`
}

// String is not required by pop and may be deleted
func (c ContactInfo) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// ContactInfos is not required by pop and may be deleted
type ContactInfos []ContactInfo

// String is not required by pop and may be deleted
func (c ContactInfos) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *ContactInfo) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: c.ParentName, Name: "ParentName"},
		&validators.StringIsPresent{Field: c.PhoneNumber, Name: "PhoneNumber"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *ContactInfo) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *ContactInfo) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
