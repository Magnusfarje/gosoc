package models

import "time"

// User model
type User struct {
	ID        string `json:"id"`
	Mail      string `json:"email"`
	Provider  string
	Picture   string
	FirstName string
	LastName  string
	Expire    time.Time
}
