package mappers

import (
	"github.com/ahmadalaik/be-invitation/internal/dto"
	"github.com/ahmadalaik/be-invitation/internal/models"
)

func ToTemplateResponse(t models.Template) dto.TemplateResponse {
	return dto.TemplateResponse{
		ID:          t.ID,
		Name:        t.Name,
		Slug:        t.Slug,
		Description: t.Description,
		Thumbnail:   t.Thumbnail,
	}
}
