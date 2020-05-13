package user

import (
	"Embassy/internal/middlewares"
	"Embassy/internal/middlewares/validations"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func Routes(router *mux.Router, db *gorm.DB) *mux.Router {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)
	router.Handle("/user/admin",
		negroni.New(
			negroni.HandlerFunc(middlewares.RequireAdminRights),
			negroni.HandlerFunc(validations.ReturnHandler(db).InputCreateAccount),
			negroni.HandlerFunc(handler.CreateAdmin),
			)).Methods("POST")
	router.Handle("/user/admin/login",
		negroni.New(
			negroni.HandlerFunc(validations.ReturnHandler(db).InputLogin),
			negroni.HandlerFunc(handler.Login),
		)).Methods("POST")
	router.Handle("/user/admin",
		negroni.New(
			negroni.HandlerFunc(validations.ReturnHandler(db).InputCreateAccount),
			negroni.HandlerFunc(handler.CreateAdmin),
		)).Methods("POST")
	router.Handle("/auth/google/login",
		negroni.New(negroni.HandlerFunc(handler.GoogleLogin))).Methods("GET")
	router.Handle("/auth/google/callback",
		negroni.New(negroni.HandlerFunc(handler.GoogleCallBack))).Methods("GET")
	router.Handle("/auth/facebook/login",
		negroni.New(negroni.HandlerFunc(handler.FacebookLogin))).Methods("GET")
	router.Handle("/auth/facebook/callback",
		negroni.New(negroni.HandlerFunc(handler.GoogleCallBack))).Methods("GET")
	router.Handle("/user",
		negroni.New(
			//negroni.HandlerFunc(middlewares.RequireAdminRights),
			negroni.HandlerFunc(handler.FindAll),
			)).Methods("GET")
	router.Handle("/all",
		negroni.New(
			negroni.HandlerFunc(middlewares.RequireAdminRights),
			negroni.HandlerFunc(handler.GetAll),
		)).Methods("GET")
	router.Handle("/user/{userID}",
		negroni.New(
			negroni.HandlerFunc(middlewares.RequireAdminRights),
			negroni.HandlerFunc(handler.FindByID),
			)).Methods("GET")

	return router
}
