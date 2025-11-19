package repository

import (
	"github.com/ahmadalaik/be-invitation/internal/models"
	"gorm.io/gorm"
)

type StoryImageRepository interface {
	CreateMultiple(images []models.StoryImage) error
	FindByInvitationID(invID uint) ([]models.StoryImage, error)
	DeleteByInvitationID(invID uint) error
}

type storyImageRepository struct {
	db *gorm.DB
}

func NewStoryImageRepository(db *gorm.DB) StoryImageRepository {
	return &storyImageRepository{db}
}

func (r *storyImageRepository) CreateMultiple(images []models.StoryImage) error {
	return r.db.Create(&images).Error
}

func (r *storyImageRepository) FindByInvitationID(invID uint) ([]models.StoryImage, error) {
	var storyImages []models.StoryImage
	if err := r.db.Where("invitation_id = ?", invID).Find(&storyImages).Error; err != nil {
		return nil, err
	}
	return storyImages, nil
}

func (r *storyImageRepository) DeleteByInvitationID(invID uint) error {
	return r.db.Where("invitation_id = ?", invID).Delete(&models.StoryImage{}).Error
}


