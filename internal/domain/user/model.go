package user

import (
	"embassy/internal/database"
	"embassy/internal/domain/registration"
)

type User struct {
	database.Base
	Email string `gorm:"type:varchar(100);unique_index;not_null"`
	Password string `gorm:"type:varchar(250);"`
	Name string `gorm:"type:varchar(250);"`
	IsAdmin bool `gorm:"type:boolean;default:false"`
	IsVerified bool `gorm:"type:boolean;not_null"`
	RegistrationDetails registration.Registration
}
