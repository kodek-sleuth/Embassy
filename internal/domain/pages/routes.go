package pages

import (
	"Embassy/internal/middlewares"
	"Embassy/internal/middlewares/validations"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/urfave/negroni"
)

func Routes(router *mux.Router, db *gorm.DB) *mux.Router {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)
	router.Handle("/pages",
		negroni.New(
			//negroni.HandlerFunc(middlewares.RequireTokenAuthentication),
			//negroni.HandlerFunc(middlewares.RequireAdminRights),
			negroni.HandlerFunc(validations.ReturnHandler(db).InputPages),
			negroni.HandlerFunc(handler.Create))).Methods("POST")
	router.Handle("/pages/{pageType}",
		negroni.New(
			negroni.HandlerFunc(middlewares.RequireTokenAuthentication),
			negroni.HandlerFunc(middlewares.RequireAdminRights),
			negroni.HandlerFunc(validations.ReturnHandler(db).InputPages),
			negroni.HandlerFunc(handler.Update))).Methods("PUT")
	router.Handle("/pages",
		negroni.New(negroni.HandlerFunc(handler.FindAll))).Methods("GET")
	router.Handle("/pages/{pageType}",
		negroni.New(negroni.HandlerFunc(handler.FindById))).Methods("GET")
	router.Handle("/pages/{pageType}",
		negroni.New(
			negroni.HandlerFunc(middlewares.RequireTokenAuthentication),
			negroni.HandlerFunc(middlewares.RequireAdminRights),
			negroni.HandlerFunc(handler.Delete))).Methods("DELETE")
	return router
}
