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
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// LoginAdministratorParams with()
type LoginAdministratorParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Validate valdiates the structures and returns an ValidationError type
func (a *RegisterAdministratorParams) Validate() (bool, ValidationError) {
	errors := make(ValidationError)

	if IsEmpty(a.FirstName) {
		errors["first_name"] = []string{"is required"}
	}

	if IsEmpty(a.LastName) {
		errors["last_name"] = []string{"is required"}
	}

	if IsEmpty(a.Email) {
		errors["email"] = []string{"is required"}
	}

	if !IsEmail(a.Email) {
		errors["email"] = []string{"is invalid email"}
	}

	if len(a.Password) < 6 {
		errors["password"] = []string{"should be at least 6 characters"}
	}

	if len(errors) == 0 {
		return true, nil
	}

	return false, errors
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
