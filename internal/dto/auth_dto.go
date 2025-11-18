package dto

type (
	RegisterRequest struct {
		Username string `json:"username" validate:"required,min=3"`
		Name     string `json:"name" validate:"required,min=3"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6"`
	}

	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
)

type (
	LoginResponse struct {
		AccessToken string `json:"access_token"`
	}
)
