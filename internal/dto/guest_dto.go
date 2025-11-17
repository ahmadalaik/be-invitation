package dto

type (
	CreateGuestRequest struct {
		InvitationID uint   `json:"invitation_id" validate:"required"`
		Name         string `json:"name" validate:"required"`
	}

	UpdateGuestRequest struct {
		Name        string `json:"name"`
		Email       string `json:"email"`
		IsAttending *bool  `json:"is_attending"`
		RSVPMessage string `json:"rsvp_message"`
	}
)

type (
	GuestResponse struct {
		ID           uint   `json:"id"`
		InvitationID uint   `json:"invitation_id"`
		Name         string `json:"name"`
		IsAttending  bool   `json:"is_attending"`
		RSVPMessage  string `json:"rsvp_message"`
	}
)
