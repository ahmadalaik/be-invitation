package dto

type (
	CreateInvitationRequest struct {
		TemplateID        uint64 `form:"template_id" validate:"required"`
		BrideName         string `form:"bride_name" validate:"required"`
		GroomName         string `form:"groom_name" validate:"required"`
		EventDateTime     string `form:"event_date_time" validate:"required"`
		Venue             string `form:"venue" validate:"required"`
		Address           string `form:"address" validate:"required"`
		ReceptionDateTime string `form:"reception_date_time" validate:"required"`
		ReceptionVenue    string `form:"reception_venue" validate:"required"`
		ReceptionAddress  string `form:"reception_address" validate:"required"`
		Story             string `form:"story" validate:"required"`
	}

	UpdateInvitationRequest struct {
		TemplateID        uint64 `form:"template_id" validate:"required"`
		BrideName         string `form:"bride_name" validate:"required"`
		GroomName         string `form:"groom_name" validate:"required"`
		EventDateTime     string `form:"event_date_time" validate:"required"`
		Venue             string `form:"venue" validate:"required"`
		Address           string `form:"address" validate:"required"`
		ReceptionDateTime string `form:"reception_date_time" validate:"required"`
		ReceptionVenue    string `form:"reception_venue" validate:"required"`
		ReceptionAddress  string `form:"reception_address" validate:"required"`
		Story             string `form:"story" validate:"required"`
	}
)

type (
	InvitationResponse struct {
		ID                uint64   `json:"id"`
		Slug              string   `json:"slug"`
		Hero              string   `json:"hero"`
		BrideName         string   `json:"bride_name"`
		GroomName         string   `json:"groom_name"`
		EventDateTime     string   `json:"event_date_time"`
		Venue             string   `json:"venue"`
		Address           string   `json:"address"`
		ReceptionDateTime string   `json:"reception_date_time"`
		ReceptionVenue    string   `json:"reception_venue"`
		ReceptionAddress  string   `json:"reception_address"`
		Story             string   `json:"story"`
		StoryImages       []string `json:"story_images"`
		UserID            uint64   `json:"user_id"`
		TemplateID        uint64   `json:"template_id"`
	}
)
