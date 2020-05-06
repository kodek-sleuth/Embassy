package validations

import (
	"Embassy/internal/helpers"
	"encoding/json"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gorilla/context"
	//"github.com/go-ozzo/ozzo-validation/is"
	//"github.com/sirupsen/logrus"
	"net/http"

	//"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Handler interface {
	InputCreateAccount(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
	InputLogin(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
	InputRegistration(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
	InputMessage(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
	InputNews(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
}

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &handler{
		service,
	}
}

func(s *handler) InputCreateAccount(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	err := validation.ValidateStruct(&user,
		validation.Field(&user.Name,
			validation.Required.Error("name is required"),
			validation.Length(1, 50)),
		validation.Field(&user.Email,
			validation.Required.Error("email is required"),
			is.Email.Error("please provide a valid email")),
		validation.Field(&user.Password, validation.Required.Error("password is required")),
	)

	if err != nil{
		helpers.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"errors": err,
		})
		return
	}

	context.Set(r, "user", user)

	next(w, r)

	return
}

func (s *handler) InputLogin(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	err := validation.ValidateStruct(&user,
		validation.Field(&user.Email, validation.Required.Error("email is required")),
		validation.Field(&user.Password, validation.Required.Error("password is required")),
	)

	if err != nil{
		helpers.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"errors": err,
		})
		return
	}

	context.Set(r, "user", user)

	next(w, r)

	return
}

func (s *handler) InputRegistration(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var registration Registration

	registration.Gender = r.FormValue("gender")
	registration.FirstName = r.FormValue("firstname")
	registration.Surname = r.FormValue("surname")
	registration.PassportNumber = r.FormValue("passport_number")
	registration.City = r.FormValue("city")
	registration.Address = r.FormValue("address")
	registration.Marriage = r.FormValue("marriage")
	registration.KinName = r.FormValue("kin_name")
	registration.KinContact = r.FormValue("kin_contact")
	registration.KinRelationship = r.FormValue("kin_relationship")
	registration.OriginArea = r.FormValue("origin_area")
	registration.ArrivalDate = r.FormValue("arrival_date")
	registration.Comment = r.FormValue("comment")

	files, err := helpers.FileUpload(r, []string{"proof_of_residence", "passport_photo"})
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	registration.ProofOfResidence = files["proof_of_residence"]
	registration.Photo = files["passport_photo"]

	err = validation.ValidateStruct(&registration,
		validation.Field(&registration.Gender, validation.Required.Error("please provide a gender")),
		validation.Field(&registration.FirstName, validation.Required.Error("please provide a first name")),
		validation.Field(&registration.Surname, validation.Required.Error("please provide a surname")),
		validation.Field(&registration.City, validation.Required.Error("please provide a photo")),
		validation.Field(&registration.Photo, validation.Required.Error("please provide passport photo")),
		validation.Field(&registration.ProofOfResidence, validation.Required.Error("please provide proof of residence")),
		validation.Field(&registration.Address, validation.Required.Error("please provide address")),
		validation.Field(&registration.OriginArea, validation.Required.Error("please provide area of origin")),
		validation.Field(&registration.KinContact, validation.Required.Error("please provide the kin's contact")),
		validation.Field(&registration.KinName, validation.Required.Error("please provide kin's name")),
		validation.Field(&registration.KinRelationship, validation.Required.Error("please provide relationship you have with kin")),
		validation.Field(&registration.ArrivalDate, validation.Required.Error("please provide your date of arrival")),
		validation.Field(&registration.Marriage, validation.Required.Error("please provide your marriage status")),
	)

	if err != nil{
		helpers.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"errors": err,
		})
		return
	}

	next(w, r)

	return
}


func (s *handler) InputNews(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var news News

	news.Title = r.FormValue("title")
	news.Body = r.FormValue("body")

	files, err := helpers.FileUpload(r, []string{"cover"})
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	news.Image = files["cover"]

	err = validation.ValidateStruct(&news,
		validation.Field(&news.Title, validation.Required.Error("title is required")),
		validation.Field(&news.Body, validation.Required.Error("description is required")),
		validation.Field(&news.Image, validation.Required.Error("cover photo is required")),
	)

	if err != nil{
		helpers.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"errors": err,
		})
		return
	}

	context.Set(r, "news", news)

	next(w, r)

	return
}

func (s *handler) InputMessage(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	panic("implement me")
}

