package tourism

import (
	"Embassy/internal/helpers"
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
	var tourism Tourism

	tourism.Title = r.FormValue("title")
	tourism.Body = r.FormValue("body")

	files, err := helpers.FileUpload(r, []string{"cover"})
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	tourism.Image = files["cover"]

	userDetails, _ := helpers.VerifyToken(r)
	tourism.UserID = userDetails.ID

	result, err := s.service.Create(&tourism)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, result)
	return
}

func (s *handler) Update(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	var tourism Tourism
	tourismID := mux.Vars(r)["tourismID"]

	parsedTourismID, err := uuid.FromString(tourismID)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	tourism.Title = r.FormValue("title")
	tourism.Body = r.FormValue("body")

	files, err := helpers.FileUpload(r, []string{"cover"})
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	tourism.Image = files["cover"]

	tourism.ID = parsedTourismID

	entity, err := s.service.Update(&tourism)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusAccepted, entity)
	return
}

func (s *handler) Delete(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	var tourism Tourism
	tourismID := mux.Vars(r)["tourismID"]

	parsedTourismID, err := uuid.FromString(tourismID)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	tourism.ID = parsedTourismID

	if err = s.service.Delete(&tourism); err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusNoContent, map[string]string{
		"message": "tourism deleted successfully",
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
	var tourism Tourism
	tourismID := mux.Vars(r)["tourismID"]
	parsedID, err := uuid.FromString(tourismID)

	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	tourism.ID = parsedID
	result, err := s.service.FindById(&tourism)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, result)
	return
}

