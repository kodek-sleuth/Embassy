package chatting

import (
	"Embassy/internal/middlewares"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/urfave/negroni"
)

func Routes(router *mux.Router, db *gorm.DB) *mux.Router {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)
	router.Handle("/chatting/{userID}",
		negroni.New(
			//negroni.HandlerFunc(middlewares.RequireTokenAuthentication),
			//negroni.HandlerFunc(middlewares.RequireAdminRights),
			negroni.HandlerFunc(handler.Chatting)))
	router.Handle("/chatting/{chattingID}",
		negroni.New(
			negroni.HandlerFunc(middlewares.RequireTokenAuthentication),
			negroni.HandlerFunc(middlewares.RequireAdminRights),
			negroni.HandlerFunc(handler.Update))).Methods("PUT")
	router.Handle("/chatting",
		negroni.New(negroni.HandlerFunc(handler.FindAll))).Methods("GET")
	router.Handle("/chat/{requestedUserID}",
		negroni.New(
			negroni.HandlerFunc(middlewares.RequireTokenAuthentication),
			negroni.HandlerFunc(middlewares.RequireAdminRights),
			negroni.HandlerFunc(handler.FindByID))).Methods("GET")
	router.Handle("/chatting/{chattingID}",
		negroni.New(
			negroni.HandlerFunc(middlewares.RequireTokenAuthentication),
			negroni.HandlerFunc(handler.Delete))).Methods("DELETE")
	return router
}
