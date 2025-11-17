package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar;size:100;not null"`
	Email     string `gorm:"type:varchar;size:100;uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Invitations []Invitation `gorm:"foreignKey:UserID"`
}
