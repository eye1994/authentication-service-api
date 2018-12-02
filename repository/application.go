package repository

import "time"

// Application with()
type Application struct {
	ID        uint       `gorm:"primary_key; AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`

	Name        string `gorm:"not null"`
	URL         string `gorm:"not null"`
	PublicToken string `gorm:"not null"`
	SecretToken string `gorm:"not null"`

	Administrators []*Administrator `gorm:"many2many:application_administrators;"`
	Users          []*User          `gorm:"many2many:application_users;"`
}
