package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"strings"
)

// Credential represents a user of system
type Credential struct {
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
}

// ValidateAndNormalizeCredential validates and normalize credentials data
func (c Credential) ValidateAndNormalizeCredential() error {

	if c.Password == "" {
		return errors.New("Invalid password!")
	}

	if c.Email == "" {
		return errors.New("Invalid email!")
	}

	if error := checkmail.ValidateFormat(c.Email); error != nil {
		return errors.New("Invalid email format!")
	}

	c.Email = strings.TrimSpace(c.Email)
	c.Email = strings.ToLower(c.Email)

	return nil
}