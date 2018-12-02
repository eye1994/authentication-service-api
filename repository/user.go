package repository

import "time"

// User User with()
type User struct {
	ID        uint       `gorm:"primary_key; AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`

	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Email     string `gorm:"type:varchar(100);unique_index" json:"email"`
	Password  string `gorm:"not null" json:"-"`

	Applications []*Application `gorm:"many2many:application_users;"`
}
