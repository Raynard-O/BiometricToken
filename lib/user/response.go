package user

import "time"

type UserResonse struct {
	FullName string `json:"full_name"`
	Email string `json:"email"`
	BioAuth bool	`json:"bio_auth"`
	VerifiedAt time.Time `json:"verified_at"`
	Active    bool   `json:"active"`
}

