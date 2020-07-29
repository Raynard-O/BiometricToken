package admin

import "time"

type AdminLoginResponse struct {
	FullName string `json:"full_name"`
	Email string `json:"email"`
	BioAuth bool	`json:"bio_auth"`
	Token	string	`json:"token"`
	VerifiedAt time.Time `json:"verified_at"`
	Active    bool   `json:"active"`
}
