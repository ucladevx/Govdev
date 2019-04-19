package models

import "time"

// User is a struct that houses users
type User struct {
	Userid    string
	Username  string
	Email     string
	Password  string
	FirstName string
	LastName  string

	// TODO Permissions

	CreatedAt time.Time
	DeletedAt time.Time
}
