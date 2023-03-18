package database

import "time"

type User struct {
	IdentityNumber string
	Password       string
	Name           string
	Role           int
	Gender         int
	Age            int
	PhoneNumber    string
	LastLogin      time.Time
	Avatar         string
	Email          string
}

func (User) TableName() string {
	return "user"
}
