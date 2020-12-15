package models

import "time"

// represents a user of system
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"string,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
