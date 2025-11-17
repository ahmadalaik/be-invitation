package models

import "time"

type Template struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar;size:150;not null"`
	Slug        string `gorm:"type:varchar;size:150;uniqueIndex;not null"`
	Description string `gorm:"type:text"`
	Thumbnail   string `gorm:"type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Invitations []Invitation `gorm:"foreignKey:TemplateID"`
}
