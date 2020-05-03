package news

import (
	"embassy/internal/database"
	uuid "github.com/satori/go.uuid"
)

type News struct {
	database.Base
	Image string `gorm:"type:varchar(100);not_null"`
	Title string `gorm:"type:varchar(100);not_null"`
	Body string `gorm:"type:varchar(100);not_null"`
	UserID uuid.UUID `gorm:"type:uuid;not_null"`
}
