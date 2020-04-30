package news

import (
	"Embassy/internal/helpers"
	"encoding/json"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
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
	var news News

	if err := json.NewDecoder(r.Body).Decode(&news); err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	userDetails, _ := helpers.VerifyToken(r)
	news.UserID = userDetails.ID

	result, err := s.service.Create(&news)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, result)
	return
}

func (s *handler) Update(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	var news News
	newsID := mux.Vars(r)["newsID"]

	parsedNewsID, err := uuid.FromString(newsID)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&news); err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	news.ID = parsedNewsID

	result, err := s.service.Update(&news)
	if err != nil{
		if err.Error() == "is not owner" {
			helpers.ErrorResponse(w, http.StatusForbidden,
				"failed to perform action, please contact administration for help")
			return
		}
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusAccepted, result)
	return
}

func (s *handler) Delete(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	var news News
	newsID := mux.Vars(r)["newsID"]

	parsedNewsID, err := uuid.FromString(newsID)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	news.ID = parsedNewsID

	if err = s.service.Delete(&news); err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusNoContent, map[string]string{
		"message": "news deleted successfully",
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
	var news News
	newsID := mux.Vars(r)["newsID"]
	parsedID, err := uuid.FromString(newsID)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	news.ID = parsedID
	result, err := s.service.FindById(&news)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, result)
	return
}

