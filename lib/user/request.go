package user

type RegisterParams struct {
	FullName string `json:"full_name"`
	Email string `json:"email"`
	Password string `json:"password"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
	BioAuth bool	`json:"bio_auth"`
}

type VerifyParams struct {
	Email	string	`json:"email"`
	BioAuth	bool	`json:"bio_auth"`
}





