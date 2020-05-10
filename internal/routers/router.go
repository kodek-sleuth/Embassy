package routers

import (
	"Embassy/internal/domain/chatting"
	"Embassy/internal/domain/education"
	"Embassy/internal/domain/pages"
	"Embassy/internal/domain/notice"
	"Embassy/internal/domain/registration"
	"Embassy/internal/domain/tourism"
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
		&pages.News{},
		&notice.Notice{},
		&education.Education{},
		&tourism.Tourism{},
		)

	router := mux.NewRouter()
	newRouter := router.PathPrefix("/api").Subrouter()
	newRouter = user.Routes(newRouter, db)
	newRouter = registration.Routes(newRouter, db)
	newRouter = chatting.Routes(newRouter, db)
	newRouter = pages.Routes(newRouter, db)
	newRouter = notice.Routes(newRouter, db)
	newRouter = education.Routes(newRouter, db)
	newRouter = tourism.Routes(newRouter, db)
	return newRouter
}
