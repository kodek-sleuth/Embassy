package validations

import (
	"github.com/jinzhu/gorm"
)

func ReturnHandler(db *gorm.DB) Handler {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)
	return handler
}
