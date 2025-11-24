package service

import (
	"fmt"
	"time"

	"github.com/ahmadalaik/be-invitation/internal/dto"
	"github.com/ahmadalaik/be-invitation/internal/mappers"
	"github.com/ahmadalaik/be-invitation/internal/models"
	"github.com/ahmadalaik/be-invitation/internal/repository"
	"github.com/ahmadalaik/be-invitation/pkg/slug"
)

type InvitationService interface {
	CreateInvitation(userID uint64, req dto.CreateInvitationRequest, mainImageURL string, storyImageURLs []string) (*dto.InvitationResponse, error)
	GetInvitationBySlug(invSlug string) (*dto.InvitationResponse, error)
	UpdateInvitationBySlug(invSlug string, req dto.UpdateInvitationRequest, newMainImageURL string) error
	GetAllInvitationsByUserID(userID uint64) ([]dto.InvitationResponse, error)
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

func (s *invitationService) CreateInvitation(userID uint64, req dto.CreateInvitationRequest, mainImageURL string, storyImageURLs []string) (*dto.InvitationResponse, error) {
	eventDateTime, err := time.Parse(time.RFC3339, req.EventDateTime)
	if err != nil {
		return nil, err
	}

	receptionDateTime, err := time.Parse(time.RFC3339, req.ReceptionDateTime)
	if err != nil {
		return nil, err
	}

	slug := fmt.Sprintf("%s-and-%s", slug.GenerateSlug(req.BrideName), slug.GenerateSlug(req.GroomName))

	invitation := models.Invitation{
		UserID:            userID,
		TemplateID:        req.TemplateID,
		Slug:              slug,
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

func (s *invitationService) GetInvitationBySlug(invSlug string) (*dto.InvitationResponse, error) {
	invitation, err := s.invRepo.FindBySlug(invSlug)
	if err != nil {
		return nil, err
	}

	response := mappers.ToInvitationResponse(*invitation)
	return &response, nil
}

func (s *invitationService) UpdateInvitationBySlug(invSlug string, req dto.UpdateInvitationRequest, newMainImageURL string) error {
	invitation, err := s.invRepo.FindBySlug(invSlug)
	if err != nil {
		return err
	}

	newInvSlug := fmt.Sprintf("%s-and-%s", slug.GenerateSlug(req.BrideName), slug.GenerateSlug(req.GroomName))

	eventDateTime, err := time.Parse(time.RFC3339, req.EventDateTime)
	if err != nil {
		return err
	}

	receptionDateTime, err := time.Parse(time.RFC3339, req.ReceptionDateTime)
	if err != nil {
		return err
	}

	invitation.TemplateID = req.TemplateID
	invitation.Slug = newInvSlug
	invitation.Hero = newMainImageURL
	invitation.BrideName = req.BrideName
	invitation.GroomName = req.GroomName
	invitation.EventDateTime = eventDateTime
	invitation.Venue = req.Venue
	invitation.Address = req.Address
	invitation.ReceptionDateTime = receptionDateTime
	invitation.ReceptionVenue = req.ReceptionVenue
	invitation.ReceptionAddress = req.ReceptionAddress
	invitation.Story = req.Story

	return s.invRepo.Update(invitation)
}

func (s *invitationService) GetAllInvitationsByUserID(userID uint64) ([]dto.InvitationResponse, error) {
	invitations, err := s.invRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	var responses []dto.InvitationResponse
	for _, inv := range invitations {
		res := mappers.ToInvitationResponse(inv)
		responses = append(responses, res)
	}

	return responses, nil
}
