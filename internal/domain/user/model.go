package user

import (
	"Embassy/internal/database"
	"Embassy/internal/domain/education"
	"Embassy/internal/domain/pages"
	"Embassy/internal/domain/notice"
	"Embassy/internal/domain/registration"
	"Embassy/internal/domain/tourism"
)

type User struct {
	database.Base
	Image string
	Email string `gorm:"type:varchar(100);unique_index;not_null"json:"email"`
	Password string `gorm:"type:varchar(250);"json:"password"`
	Name string `json:"name"`
	IsAdmin bool `gorm:"type:boolean;default:false"`
	IsVerified bool `gorm:"type:boolean;not_null"`
	RegistrationDetails registration.Registration
	Notice []notice.Notice
	Education []education.Education
	Tourism []tourism.Tourism
	News []pages.News
}
