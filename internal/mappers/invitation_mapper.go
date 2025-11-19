package mappers

import (
	"time"

	"github.com/ahmadalaik/be-invitation/internal/dto"
	"github.com/ahmadalaik/be-invitation/internal/models"
)

func ToInvitationResponse(inv models.Invitation) dto.InvitationResponse {
	var storyImageURLs []string

	for _, storyImage := range inv.StoryImages {
		storyImageURLs = append(storyImageURLs, storyImage.ImageURL)
	}

	return dto.InvitationResponse{
		ID:                inv.ID,
		Slug:              inv.Slug,
		Hero:              inv.Hero,
		BrideName:         inv.BrideName,
		GroomName:         inv.GroomName,
		EventDateTime:     inv.EventDateTime.Format(time.RFC3339),
		Venue:             inv.Venue,
		Address:           inv.Address,
		ReceptionDateTime: inv.ReceptionDateTime.Format(time.RFC3339),
		ReceptionVenue:    inv.ReceptionVenue,
		ReceptionAddress:  inv.ReceptionAddress,
		Story:             inv.Story,
		StoryImages:       storyImageURLs,
		UserID:            inv.UserID,
		TemplateID:        inv.TemplateID,
	}
}
