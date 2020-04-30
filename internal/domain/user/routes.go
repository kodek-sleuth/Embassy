package user

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func Routes(router *mux.Router, db *gorm.DB) *mux.Router {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)
	router.Handle("/user/signup",
		negroni.New(negroni.HandlerFunc(handler.Create))).Methods("POST")
	router.Handle("/auth/google/login",
		negroni.New(negroni.HandlerFunc(handler.GoogleLogin))).Methods("GET")
	router.Handle("/auth/google/callback",
		negroni.New(negroni.HandlerFunc(handler.GoogleCallBack))).Methods("GET")
	router.Handle("/auth/facebook/login",
		negroni.New(negroni.HandlerFunc(handler.FacebookLogin))).Methods("GET")
	router.Handle("/auth/facebook/callback",
		negroni.New(negroni.HandlerFunc(handler.GoogleCallBack))).Methods("GET")
	router.Handle("/user",
		negroni.New(negroni.HandlerFunc(handler.FindAll))).Methods("GET")
	router.Handle("/user/{userID}",
		negroni.New(negroni.HandlerFunc(handler.FindByID))).Methods("GET")

	return router
}
