package user

type RegisterParams struct {
	FullName string `json:"full_name"`
	Email string `json:"email"`
	Password string `json:"password"`
	BioAuth bool	`json:"bio_auth"`
}


