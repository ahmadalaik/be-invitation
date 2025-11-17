package mappers

import (
	"github.com/ahmadalaik/be-invitation/internal/dto"
	"github.com/ahmadalaik/be-invitation/internal/models"
)

func ToGuestResponse(g models.Guest) dto.GuestResponse {
	return dto.GuestResponse{
		ID:           g.ID,
		InvitationID: g.InvitationID,
		Name:         g.Name,
		IsAttending:  g.IsAttending,
		RSVPMessage:  g.RSVPMessage,
	}
}
