package service

import (
	"github.com/ahmadalaik/be-invitation/internal/dto"
	"github.com/ahmadalaik/be-invitation/internal/models"
	"github.com/ahmadalaik/be-invitation/internal/repository"
)

type TemplateService interface {
	CreateTemplate(req dto.CreateTemplateRequest) error
}

type templateService struct {
	templateRepo repository.TemplateRepository
}

func NewTemplateService(templateRepo repository.TemplateRepository) TemplateService {
	return &templateService{templateRepo}
}

func (s *templateService) CreateTemplate(req dto.CreateTemplateRequest) error {
	template := models.Template{
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
		Thumbnail:   req.Thumbnail,
	}

	return s.templateRepo.Create(template)
}
