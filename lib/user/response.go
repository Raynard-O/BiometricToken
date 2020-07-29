package user

import "time"

type UserVerifyResonse struct {
	FullName string `json:"full_name"`
	Email string `json:"email"`
	BioAuth bool	`json:"bio_auth"`
	VerifiedAt time.Time `json:"verified_at"`
	Active    bool   `json:"active"`
}




type UserResponse struct {
	FullName string `json:"full_name"`
	Email string `json:"email"`
	BioAuth bool	`json:"bio_auth"`
	VerifiedAt time.Time `json:"verified_at"`
	Active    bool   `json:"active"`
}

type AdminDetail struct {
	AdminName	string	`json:"admin_name"`
	AdminID		uint	`json:"admin_id"`
}

type RegisterResponse struct {
	User UserResponse	`json:"user"`
	Success bool	`json:"success"`
	Admin	AdminDetail	`json:"admin"`
}

