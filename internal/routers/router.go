package routers

import (
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
		)

	router := mux.NewRouter()
	newRouter := router.PathPrefix("/api").Subrouter()
	newRouter = user.Routes(newRouter, db)
	newRouter = registration.Routes(newRouter, db)
	return newRouter
}
