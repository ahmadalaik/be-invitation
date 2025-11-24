package dto

type (
	CreateTemplateRequest struct {
		Name        string `json:"name" validate:"required,min=3"`
		Slug        string `json:"slug" validate:"required"`
		Description string `json:"description"`
		Thumbnail   string `json:"thumbnail"`
	}

	UpdateTemplateRequest struct {
		Name        string `json:"name" validate:"required,min=3"`
		Description string `json:"description" validate:"omitempty"`
		Thumbnail   string `json:"thumbnail" validate:"omitempty"`
	}
)

type (
	TemplateResponse struct {
		ID          uint64 `json:"id"`
		Name        string `json:"name"`
		Slug        string `json:"slug"`
		Description string `json:"description"`
		Thumbnail   string `json:"thumbnail"`
	}
)
