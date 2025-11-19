package repository

import (
	"github.com/ahmadalaik/be-invitation/internal/models"
	"gorm.io/gorm"
)

type InvitationRepository interface {
	Create(inv *models.Invitation) error
	FindByID(id uint) (*models.Invitation, error)
	FindBySlug(slug string) (*models.Invitation, error)
	FindByUserID(userID uint) ([]models.Invitation, error)
	Update(inv models.Invitation) error
	Delete(id uint) error
}

type invitationRepository struct {
	db *gorm.DB
}

func NewInvitationRepository(db *gorm.DB) InvitationRepository {
	return &invitationRepository{db}
}

// Create implements InvitationRepository.
func (r *invitationRepository) Create(inv *models.Invitation) error {
	return r.db.Create(inv).Error
}

// FindByID implements InvitationRepository.
func (r *invitationRepository) FindByID(id uint) (*models.Invitation, error) {
	var inv models.Invitation
	if err := r.db.Preload("Guests").Preload("Template").Where("id = ?", id).First(&inv).Error; err != nil {
		return nil, err
	}
	return &inv, nil
}

func (r *invitationRepository) FindBySlug(slug string) (*models.Invitation, error) {
	var inv models.Invitation
	if err := r.db.Preload("Guests").Preload("Template").Where("slug = ?", slug).First(&inv).Error; err != nil {
		return nil, err
	}
	return &inv, nil
}

// FindByUserID implements InvitationRepository.
func (r *invitationRepository) FindByUserID(userID uint) ([]models.Invitation, error) {
	var invs []models.Invitation
	if err := r.db.Where("user_id = ?", userID).Find(&invs).Error; err != nil {
		return nil, err
	}
	return invs, nil
}

// Update implements InvitationRepository.
func (r *invitationRepository) Update(inv models.Invitation) error {
	return r.db.Save(&inv).Error
}

// Delete implements InvitationRepository.
func (r *invitationRepository) Delete(id uint) error {
	var inv models.Invitation
	return r.db.Where("id = ?", id).Delete(&inv).Error
}
