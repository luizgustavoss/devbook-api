package models

// PasswordUpdate represents data for password update
type PasswordUpdate struct {
	PreviousPassword string `json:"previous,omitempty"`
	NewPassword      string `json:"new,omitempty"`
}
