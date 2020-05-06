package registration

import (
	"Embassy/internal/helpers"
	"github.com/sirupsen/logrus"
	"net/http"
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

	auth, err := helpers.VerifyToken(r)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, "failed to authenticate user")
		return
	}

	registration.UserID = auth.ID

	user, err := u.service.Create(&registration)
	if err != nil{
		logrus.Println(err)
		helpers.ErrorResponse(w, http.StatusInternalServerError, "failed to register user")
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, map[string]interface{}{
		"message": "you have successfully registered",
		"code": user.Code,
	})
	return
}

func (u *handler) Update(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
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

	auth, err := helpers.VerifyToken(r)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, "failed to authenticate user")
		return
	}

	registration.UserID = auth.ID

	_, err = u.service.Update(&registration)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, "failed to update user details")
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, map[string]interface{}{
		"message": "you have successfully updated your details",
	})
	return
}
