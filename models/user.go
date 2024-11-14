package models

import "time"

type User struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Gender    string    `json:"gender" gorm:"not null"`
	BirthDate time.Time `json:"birthDate" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null;unique"`
	Password  string    `json:"-"`
	Role      string    `json:"role" gorm:"default:attendee"`
	Level     int       `json:"level" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
