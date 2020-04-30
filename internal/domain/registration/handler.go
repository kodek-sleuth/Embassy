package registration

import (
	"Embassy/internal/helpers"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
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
	var user Registration
	rand.Seed(time.Now().UnixNano())

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	auth, err := helpers.VerifyToken(r)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, "failed to authenticate user")
		return
	}

	user.UserID = auth.ID

	_, err = u.service.Create(&user)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, "failed to register user")
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, map[string]interface{}{
		"message": "you have successfully registered, the registration code has been sent on your email",
	})
	return
}

func (u *handler) Update(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	var user Registration

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	auth, err := helpers.VerifyToken(r)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, "failed to authenticate user")
		return
	}

	user.UserID = auth.ID

	_, err = u.service.Update(&user)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, "failed to update user details")
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, map[string]interface{}{
		"message": "you have successfully updated your details",
	})
	return
}
