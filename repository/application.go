package repository

import "time"

// Application with()
type Application struct {
	ID        uint       `gorm:"primary_key; AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`

	Name             string `gorm:"not null" json:"name"`
	URL              string `gorm:"not null" json:"url"`
	PublicToken      string `gorm:"not null" json:"public_token"`
	SecretToken      string `gorm:"not null" json:"-"`
	OpenRegistration bool   `json:"open_registration"`

	Administrators []*Administrator `gorm:"many2many:application_administrators;" json:"administrators"`
	Users          []*User          `gorm:"many2many:application_users;" json:"users"`
}

// ApplicationParams with()
type ApplicationParams struct {
	Name             string `json:"name" validate:"nonzero"`
	URL              string `json:"url" validate:"nonzero"`
	OpenRegistration bool   `json:"open_registration"`
}
