package dto

type (
	CreateInvitationRequest struct {
		UserID            uint     `json:"user_id" validate:"required"`
		TemplateID        uint     `json:"template_id" validate:"required"`
		Slug              string   `json:"slug" validate:"required"`
		Hero              string   `json:"hero" validate:"required"`
		BrideName         string   `json:"bride_name" validate:"required"`
		GroomName         string   `json:"groom_name" validate:"required"`
		EventDateTime     string   `json:"event_date_time" validate:"required"`
		Venue             string   `json:"venue" validate:"required"`
		Address           string   `json:"address" validate:"required"`
		ReceptionDateTime string   `json:"reception_date_time" validate:"required"`
		ReceptionVenue    string   `json:"reception_venue" validate:"required"`
		ReceptionAddress  string   `json:"reception_address" validate:"required"`
		Story             string   `json:"story" validate:"required"`
		StoryImages       []string `json:"story_images" validate:"required"`
	}

	UpdateInvitationRequest struct {
		Hero              string   `json:"hero" validate:"required"`
		BrideName         string   `json:"bride_name" validate:"required"`
		GroomName         string   `json:"groom_name" validate:"required"`
		EventDateTime     string   `json:"event_date_time" validate:"required"`
		Venue             string   `json:"venue" validate:"required"`
		Address           string   `json:"address" validate:"required"`
		ReceptionDateTime string   `json:"reception_date_time" validate:"required"`
		ReceptionVenue    string   `json:"reception_venue" validate:"required"`
		ReceptionAddress  string   `json:"reception_address" validate:"required"`
		Story             string   `json:"story" validate:"required"`
		StoryImages       []string `json:"story_images" validate:"required"`
	}
)

type (
	InvitationResponse struct {
		ID                uint     `json:"id"`
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
		UserID            uint     `json:"user_id"`
		TemplateID        uint     `json:"template_id"`
	}
)
