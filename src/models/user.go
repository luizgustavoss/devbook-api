package models

import (
	"errors"
	"strings"
	"time"
)

// represents a user of system
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// PrepareCreate Prepare validates and formats user data for creation
func (user *User) PrepareCreate() error {

	if error := user.validateCreation(); error != nil {
		return error
	}
	user.format()
	return nil
}

// PrepareUpdate Prepare validates and formats user data for update
func (user *User) PrepareUpdate() error {

	if error := user.validateUpdate(); error != nil {
		return error
	}
	user.format()
	return nil
}

func (user *User) validateUpdate() error {

	if error := user.validateCommonAttributes(); error != nil {
		return error
	}

	return nil
}

func (user *User) validateCreation() error {

	if error := user.validateCommonAttributes(); error != nil {
		return error
	}

	if user.Password == "" {
		return errors.New("Invalid user password!")
	}

	return nil
}

func (user *User) validateCommonAttributes() error {

	if user.Name == "" {
		return errors.New("Invalid user name!")
	}

	if user.Email == "" {
		return errors.New("Invalid user email!")
	}

	if user.Nick == "" {
		return errors.New("Invalid user nick!")
	}
	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Nick = strings.TrimSpace(user.Nick)
}