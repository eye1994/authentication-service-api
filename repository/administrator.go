package repository

import (
	"time"

	"github.com/eye1994/authentication-service-api/utils"
)

// Administrator with()
type Administrator struct {
	ID        uint       `gorm:"primary_key; AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`

	FirstName    string `gorm:"not null" json:"first_name"`
	LastName     string `gorm:"not null" json:"last_name"`
	Email        string `gorm:"type:varchar(100);unique_index" json:"email"`
	PasswordHash string `gorm:"not null" json:"-"`

	Applications []*Application `gorm:"many2many:application_administrators;" json:"-"`
}

// RegisterAdministratorParams with()
type RegisterAdministratorParams struct {
	FirstName string `json:"first_name" validate:"nonzero"`
	LastName  string `json:"last_name" validate:"nonzero"`
	Email     string `json:"email" validate:"nonzero"`
	Password  string `json:"password" validate:"min=6"`
}

// LoginAdministratorParams with()
type LoginAdministratorParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ToModel with()
func (a *RegisterAdministratorParams) ToModel() (*Administrator, error) {
	password, err := utils.GeneratePassword(a.Password)
	if err != nil {
		return nil, err
	}

	return &Administrator{
		Email:        a.Email,
		FirstName:    a.FirstName,
		LastName:     a.LastName,
		PasswordHash: string(password),
	}, nil
}
