package routers

import (
	"embassy/internal/domain/chatting"
	"embassy/internal/domain/registration"
	"embassy/internal/domain/user"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func InitRoutes(db *gorm.DB) *mux.Router {
	// Migrations
	db.AutoMigrate(
		&user.User{},
		&registration.Registration{},
		&chatting.Chat{},
		)

	router := mux.NewRouter()
	newRouter := router.PathPrefix("/api").Subrouter()
	newRouter = user.Routes(newRouter, db)
	newRouter = registration.Routes(newRouter, db)
	newRouter = chatting.Routes(newRouter, db)
	return newRouter
}
