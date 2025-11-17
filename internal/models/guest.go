package models

import "time"

type Guest struct {
	ID           uint   `gorm:"primaryKey"`
	InvitationID uint   `gorm:"not null"`
	Name         string `gorm:"type:varchar;size:150;not null"`
	IsAttending  bool   `gorm:"default:false"`
	RSVPMessage  string `gorm:"type:text"`
	CreatedAt    time.Time
	UpdatedAt    time.Time

	Invitation Invitation `gorm:"constraint:OnDelete:CASCADE"`
}
