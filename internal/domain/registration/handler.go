package registration

import (
	"Embassy/internal/helpers"
	"github.com/gorilla/context"
	"net/http"
	"reflect"
)

// Methods to be consumed by handler
type Handler interface {
	Create(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	Update(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
}

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &handler{
		service,
	}
}

func (u *handler) Create(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	var registration Registration

	result := context.Get(r, "user")
	usr := reflect.ValueOf(result)

	registration.Gender = usr.FieldByName("Gender").String()
	registration.FirstName = usr.FieldByName("FirstName").String()
	registration.Surname = usr.FieldByName("Surname").String()
	registration.PassportNumber = usr.FieldByName("PassportNumber").String()
	registration.City = usr.FieldByName("City").String()
	registration.Address = usr.FieldByName("Address").String()
	registration.Marriage = usr.FieldByName("Marriage").String()
	registration.KinName = usr.FieldByName("KinName").String()
	registration.KinContact = usr.FieldByName("KinContact").String()
	registration.KinRelationship = usr.FieldByName("KinRelationship").String()
	registration.OriginArea = usr.FieldByName("OriginArea").String()
	registration.ArrivalDate = usr.FieldByName("ArrivalDate").String()
	registration.Comment = usr.FieldByName("Comment").String()
	registration.ProofOfResidence = usr.FieldByName("ProofOfResidence").String()
	registration.Photo = usr.FieldByName("Photo").String()

	auth, err := helpers.VerifyToken(r)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, "failed to authenticate user")
		return
	}

	registration.UserID = auth.ID

	user, err := u.service.Create(&registration)
	if err != nil{
		if err.Error() == "you already registered, this is your code CON-AUS-PLE-1664585"{
			helpers.ErrorResponse(w, http.StatusConflict,
				"you already registered, this is your code CON-AUS-PLE-1664585")
			return
		}
	}

	helpers.JSONResponse(w, http.StatusCreated, map[string]interface{}{
		"message": "you have successfully registered",
		"code": user.Code,
	})
	return
}

func (u *handler) Update(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	var registration Registration

	result := context.Get(r, "user")
	usr := reflect.ValueOf(result)

	registration.Gender = usr.FieldByName("Gender").String()
	registration.FirstName = usr.FieldByName("FirstName").String()
	registration.Surname = usr.FieldByName("Surname").String()
	registration.PassportNumber = usr.FieldByName("PassportNumber").String()
	registration.City = usr.FieldByName("City").String()
	registration.Address = usr.FieldByName("Address").String()
	registration.Marriage = usr.FieldByName("Marriage").String()
	registration.KinName = usr.FieldByName("KinName").String()
	registration.KinContact = usr.FieldByName("KinContact").String()
	registration.KinRelationship = usr.FieldByName("KinRelationship").String()
	registration.OriginArea = usr.FieldByName("OriginArea").String()
	registration.ArrivalDate = usr.FieldByName("ArrivalDate").String()
	registration.Comment = usr.FieldByName("Comment").String()
	registration.ProofOfResidence = usr.FieldByName("ProofOfResidence").String()
	registration.Photo = usr.FieldByName("Photo").String()

	auth, err := helpers.VerifyToken(r)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, "failed to authenticate user")
		return
	}

	registration.UserID = auth.ID

	_, err = u.service.Create(&registration)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, "failed to update user")
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, map[string]interface{}{
		"message": "you have successfully updated your details",
	})
	return
}
