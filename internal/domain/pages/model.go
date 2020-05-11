package pages

import (
	"Embassy/internal/database"
	uuid "github.com/satori/go.uuid"
)

type Pages struct {
	database.Base
	Type string `gorm:"type:varchar;not_null;unique"`
	Title string `gorm:"type:varchar;not_null"`
	Body string `gorm:"type:text;not_null"`
	UserID uuid.UUID `gorm:"type:uuid;not_null"`
}
