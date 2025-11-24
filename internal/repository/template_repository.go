package repository

import (
	"github.com/ahmadalaik/be-invitation/internal/models"
	"gorm.io/gorm"
)

type TemplateRepository interface {
	Create(template models.Template) error
	FindByID(id uint) (*models.Template, error)
	Update(template models.Template) error
	Delete(id uint) error
	FindAll() ([]models.Template, error)
}

type templateRepository struct {
	db *gorm.DB
}

func NewTemplateRepository(db *gorm.DB) TemplateRepository {
	return &templateRepository{db}
}

func (r *templateRepository) Create(template models.Template) error {
	return r.db.Create(&template).Error
}

func (r *templateRepository) FindByID(id uint) (*models.Template, error) {
	var template models.Template
	if err := r.db.Where("id = ?", id).First(&template).Error; err != nil {
		return nil, err
	}
	return &template, nil
}

func (r *templateRepository) Update(template models.Template) error {
	return r.db.Save(&template).Error
}

func (r *templateRepository) Delete(id uint) error {
	var template models.Template
	return r.db.Where("id = ?", id).Delete(&template).Error
}

func (r *templateRepository) FindAll() ([]models.Template, error) {
	var templates []models.Template
	if err := r.db.Find(&templates).Error; err != nil {
		return nil, err
	}
	return templates, nil
}
