package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string `json:"username" gorm:"unique;type:varchar(20)" validate:"required,min=1,max=32"`
	Password  string `json:"password" gorm:"type:varchar(255);" validate:"required,min=6"`
	FullName  string `json:"full_name" gorm:"type:varchar(100);" validate:"required,min=6"`
}

func (l User) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type UserSession struct {
	ID                  uint `gorm:"primarykey"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	UserID              int       `json:"username" gorm:"type:int" validate:"required"`
	Token               int       `json:"token" gorm:"type:varchar(255)" validate:"required"`
	RefreshToken        int       `json:"refresh_token" gorm:"type:varchar(255)" validate:"required"`
	TokenExpired        time.Time `json:"-" validate:"required"`
	RefreshTokenExpired time.Time `json:"-" validate:"required"`
}

func (l UserSession) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
