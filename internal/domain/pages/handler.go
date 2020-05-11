package pages

import (
	"Embassy/internal/helpers"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"reflect"

	//"github.com/sirupsen/logrus"
	//"golang.org/x/net/html"
	"net/http"
)

type Handler interface {
	Create(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	Update(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	Delete(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	FindAll(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	FindById(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
}

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &handler{
		service,
	}
}

func (s *handler) Create(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	var pages Pages

	result := context.Get(r, "pages")
	pgs := reflect.ValueOf(result)

	pages.Title = pgs.FieldByName("Title").String()
	pages.Body = pgs.FieldByName("Body").String()
	pages.Type = pgs.FieldByName("Type").String()

	userDetails, _ := helpers.VerifyToken(r)
	pages.UserID = userDetails.ID

	result, err := s.service.Create(&pages)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, result)
	return
}

func (s *handler) Update(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	var pages Pages

	result := context.Get(r, "pages")
	pgs := reflect.ValueOf(result)

	pages.Title = pgs.FieldByName("Title").String()
	pages.Body = pgs.FieldByName("Body").String()
	pages.Type = pgs.FieldByName("Type").String()

	userDetails, _ := helpers.VerifyToken(r)
	pages.UserID = userDetails.ID

	pageType := mux.Vars(r)["pageType"]
	pages.Type = pageType

	result, err := s.service.Update(&pages)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, result)
	return
}

func (s *handler) Delete(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	var pages Pages
	pageType := mux.Vars(r)["pageType"]
	pages.Type = pageType

	if err := s.service.Delete(&pages); err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusNoContent, map[string]string{
		"message": "pages deleted successfully",
	})
	return
}

func (s *handler) FindAll(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	result, err := s.service.FindAll()
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, result)
	return
}

func (s *handler) FindById(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	var pages Pages
	pageType := mux.Vars(r)["pageType"]
	pages.Type = pageType

	result, err := s.service.FindById(&pages)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, result)
	return
}

