package registration

import (
	"embassy/internal/middlewares"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func Routes(router *mux.Router, db *gorm.DB) *mux.Router {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)
	router.Handle("/user/register",
		negroni.New(
			negroni.HandlerFunc(handler.Create),
			negroni.HandlerFunc(middlewares.RequireTokenAuthentication),
			)).Methods("POST")
	router.Handle("/user/register",
		negroni.New(
			negroni.HandlerFunc(handler.Update),
			negroni.HandlerFunc(middlewares.RequireTokenAuthentication),
		)).Methods("PUT")
	return router
}
