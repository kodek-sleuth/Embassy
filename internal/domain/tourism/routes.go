package tourism

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
	router.Handle("/tourism",
		negroni.New(
			negroni.HandlerFunc(middlewares.RequireTokenAuthentication),
			negroni.HandlerFunc(middlewares.RequireAdminRights),
			negroni.HandlerFunc(validations.ReturnHandler(db).InputTourism),
			negroni.HandlerFunc(handler.Create))).Methods("POST")
	router.Handle("/tourism/{tourismID}",
		negroni.New(
			negroni.HandlerFunc(middlewares.RequireTokenAuthentication),
			negroni.HandlerFunc(middlewares.RequireAdminRights),
			negroni.HandlerFunc(validations.ReturnHandler(db).InputTourism),
			negroni.HandlerFunc(handler.Update))).Methods("PUT")
	router.Handle("/tourism",
		negroni.New(negroni.HandlerFunc(handler.FindAll))).Methods("GET")
	router.Handle("/tourism/{tourismID}",
		negroni.New(negroni.HandlerFunc(handler.FindById))).Methods("GET")
	router.Handle("/tourism/{tourismID}",
		negroni.New(
			negroni.HandlerFunc(middlewares.RequireTokenAuthentication),
			negroni.HandlerFunc(middlewares.RequireOwnerRights),
			negroni.HandlerFunc(handler.Delete))).Methods("DELETE")
	return router
}
