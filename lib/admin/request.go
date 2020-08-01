package admin

type AdminLoginParams struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	BioAuth      bool      `json:"bio_auth"`
}


type CreateAdminParams struct {
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	FullName  string `json:"full_name"`
	BioAuth      bool      `json:"bio_auth"`
}
