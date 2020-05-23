package chatting

import (
	"Embassy/internal/database"
	uuid "github.com/satori/go.uuid"
)

type Chat struct {
	database.Base
	From uuid.UUID `gorm:"type:uuid;not_null"`
	To uuid.UUID `gorm:"type:uuid;not_null"`
	Message string `gorm:"type:text;not_null"`
}
