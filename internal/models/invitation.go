package models

import "time"

type StoryImage struct {
	ID uint `gorm:"primaryKey"`

	// Relasi One-to-Many:
	// Kolom InvitationID adalah Foreign Key yang merujuk ke tabel invitations.
	// gorm:"index" membuat indeks pada kolom ini untuk mempercepat pencarian relasi.
	InvitationID uint `gorm:"not null"`

	// Kolom 'url' pada database. not null berarti kolom tidak boleh kosong.
	URL string `gorm:"type:varchar;size:255;not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Invitation struct {
	ID                uint      `gorm:"primaryKey"`
	UserID            uint      `gorm:"not null"`
	TemplateID        uint      `gorm:"not null"`
	Slug              string    `gorm:"type:varchar;size:150;uniqueIndex;not null"`
	Hero              string    `gorm:"type:text"`
	BrideName         string    `gorm:"type:varchar;size:150;not null"`
	GroomName         string    `gorm:"type:varchar;size:150;not null"`
	EventDateTime     time.Time `gorm:"not null"`
	Venue             string    `gorm:"not null"`
	Address           string    `gorm:"not null"`
	ReceptionDateTime time.Time `gorm:"not null"`
	ReceptionVenue    string    `gorm:"not null"`
	ReceptionAddress  string    `gorm:"not null"`
	Story             string    `gorm:"type:text;not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time

	User        User         `gorm:"constraint:OnDelete:CASCADE"`
	Template    Template     `gorm:"constraint:OnDelete:SET NULL"`
	StoryImages []StoryImage `gorm:"foreignKey:InvitationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Guests      []Guest      `gorm:"foreignKey:InvitationID"`
}
