package routers

import (
	"Embassy/internal/domain/chatting"
	"Embassy/internal/domain/news"
	"Embassy/internal/domain/registration"
	"Embassy/internal/domain/user"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func InitRoutes(db *gorm.DB) *mux.Router {
	// Migrations
	db.AutoMigrate(
		&user.User{},
		&registration.Registration{},
		&chatting.Chat{},
		&news.News{},
		)

	router := mux.NewRouter()
	newRouter := router.PathPrefix("/api").Subrouter()
	newRouter = user.Routes(newRouter, db)
	newRouter = registration.Routes(newRouter, db)
	newRouter = chatting.Routes(newRouter, db)
	newRouter = news.Routes(newRouter, db)
	return newRouter
}
