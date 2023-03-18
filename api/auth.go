package api

import "time"

type MeResponse struct {
	IdentityNumber string    `json:"identityNumber"`
	Name           string    `json:"name"`
	Role           int       `json:"role"`
	Gender         int       `json:"gender"`
	Age            int       `json:"age"`
	PhoneNumber    string    `json:"phoneNumber"`
	LastLogin      time.Time `json:"lastLogin"`
	Avatar         string    `json:"avatar"`
	Email          string    `json:"email"`
}
