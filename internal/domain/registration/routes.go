package registration

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
	router.Handle("/user/register",
		negroni.New(
			negroni.HandlerFunc(middlewares.RequireTokenAuthentication),
			negroni.HandlerFunc(validations.ReturnHandler(db).InputRegistration),
			negroni.HandlerFunc(handler.Create),
			)).Methods("POST")
	router.Handle("/user/register",
		negroni.New(
			negroni.HandlerFunc(middlewares.RequireTokenAuthentication),
			negroni.HandlerFunc(validations.ReturnHandler(db).InputRegistration),
			negroni.HandlerFunc(handler.Update),
		)).Methods("PUT")
	return router
}
