package service

import (
	"time"

	"github.com/ahmadalaik/be-invitation/internal/dto"
	"github.com/ahmadalaik/be-invitation/internal/mappers"
	"github.com/ahmadalaik/be-invitation/internal/models"
	"github.com/ahmadalaik/be-invitation/internal/repository"
)

type InvitationService interface {
	CreateInvitation(userID int64, req dto.CreateInvitationRequest, mainImageURL string, storyImageURLs []string) (*dto.InvitationResponse, error)
}

type invitationService struct {
	invRepo        repository.InvitationRepository
	storyImageRepo repository.StoryImageRepository
}

func NewInvitationService(invRepo repository.InvitationRepository, storyImageRepo repository.StoryImageRepository) InvitationService {
	return &invitationService{
		invRepo:        invRepo,
		storyImageRepo: storyImageRepo,
	}
}

func (s *invitationService) CreateInvitation(userID int64, req dto.CreateInvitationRequest, mainImageURL string, storyImageURLs []string) (*dto.InvitationResponse, error) {
	eventDateTime, err := time.Parse(time.RFC3339, req.EventDateTime)
	if err != nil {
		return nil, err
	}

	receptionDateTime, err := time.Parse(time.RFC3339, req.ReceptionDateTime)
	if err != nil {
		return nil, err
	}

	invitation := models.Invitation{
		UserID:            userID,
		TemplateID:        req.TemplateID,
		Slug:              req.Slug,
		Hero:              mainImageURL,
		BrideName:         req.BrideName,
		GroomName:         req.GroomName,
		EventDateTime:     eventDateTime,
		Venue:             req.Venue,
		Address:           req.Address,
		ReceptionDateTime: receptionDateTime,
		ReceptionVenue:    req.ReceptionVenue,
		ReceptionAddress:  req.ReceptionAddress,
		Story:             req.Story,
	}

	if err := s.invRepo.Create(&invitation); err != nil {
		return nil, err
	}

	if len(storyImageURLs) > 0 {
		var storyImages []models.StoryImage
		for _, url := range storyImageURLs {
			storyImages = append(storyImages, models.StoryImage{
				InvitationID: invitation.ID,
				ImageURL:     url,
			})
		}

		if err := s.storyImageRepo.CreateMultiple(storyImages); err != nil {
			return nil, err
		}

		invitation.StoryImages = storyImages
	}

	response := mappers.ToInvitationResponse(invitation)
	return &response, nil
}
