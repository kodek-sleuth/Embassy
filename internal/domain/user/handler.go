package user

import (
	"embassy/internal/helpers"
	"encoding/json"
	"fmt"
	"github.com/dchest/uniuri"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
	"net/http"
	"os"
)

// Methods to be consumed by handler
type Handler interface {
	Create(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	GoogleLogin(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	GoogleCallBack(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	FacebookLogin(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	FacebookCallBack(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	FindAll(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	FindByID(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	Delete(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	CreateAdmin(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	Login(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
}

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &handler{
		service,
	}
}

type GoogleUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
}

var googleOAuthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:5500/api/auth/google/callback",
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email",
							"https://www.googleapis.com/auth/userinfo.profile"},
	Endpoint:     google.Endpoint,
}

var faceBookOAuthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:5500/api/auth/facebook/callback",
	ClientID:     os.Getenv("FACEBOOK_CLIENT_ID"),
	ClientSecret: os.Getenv("FACEBOOK_CLIENT_SECRET"),
	Scopes:       []string{"email"},
	Endpoint:     facebook.Endpoint,
}

func (u *handler) FacebookLogin(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	oauthStateString := uniuri.New()
	url := faceBookOAuthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return
}

func (u *handler) FacebookCallBack(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	var user GoogleUser
	content, err := helpers.GetUserDataFromMedia(r, os.Getenv("FACEBOOK_URL"), faceBookOAuthConfig)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, "failed to get user from token")
		return
	}

	err = json.Unmarshal(content, &user)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusBadRequest, "failed to get user details")
		return
	}

	result, err := u.service.Create(&User{
		Name:user.Name,
		IsVerified:user.VerifiedEmail,
		Email:user.Email,
		Password: uniuri.New(),
	})
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, "failed to authenticate user")
		return
	}

	m := map[string]interface{}{
		"id": result.ID,
		"email": result.Email,
		"name": result.Name,
	}

	token, err := helpers.CreateToken(m)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, "failed to authenticate user")
		return
	}

	url := fmt.Sprintf("%s?token=%s", os.Getenv("FRONTEND_URL"), token)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return
}

func (u *handler) GoogleLogin(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	oauthStateString := uniuri.New()
	url := googleOAuthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return
}

func (u *handler) GoogleCallBack(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	var user GoogleUser
	content, err := helpers.GetUserDataFromMedia(r, os.Getenv("GOOGLE_URL"), googleOAuthConfig)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, "failed to get user from token")
		return
	}

	err = json.Unmarshal(content, &user)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusBadRequest, "failed to get user details")
		return
	}

	result, err := u.service.Create(&User{
		Name:user.Name,
		IsVerified:user.VerifiedEmail,
		Email:user.Email,
		Password: uniuri.New(),
	})

	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, "failed to authenticate user")
		return
	}

	m := map[string]interface{}{
		"id": result.ID,
		"email": result.Email,
		"name": result.Name,
	}

	token, err := helpers.CreateToken(m)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, "failed to authenticate user")
		return
	}

	url := fmt.Sprintf("%s?token=%s", os.Getenv("FRONTEND_URL"), token)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return
}

func (u *handler) Create(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	_, err := u.service.Update(&user)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, "failed to authenticate user")
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, map[string]interface{}{
		"message": "you have successfully registered, the registration code has been sent on your email",
	})
	return
}

func (u *handler) FindAll(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	users, err := u.service.FindAll()
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, "failed to fetch users")
		return
	}

	helpers.JSONResponse(w, http.StatusOK, users)
	return
}

func (u *handler) FindByID(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	var user User
	userID := mux.Vars(r)["userID"]

	id, err := helpers.ParseIDs([]string{userID})
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, "failed to fetch users")
		return
	}

	user.ID = id[0]
	result, err := u.service.FindBy(&user, "id")
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, "failed to fetch users")
		return
	}

	helpers.JSONResponse(w, http.StatusOK, result)
	return
}

func (u *handler) Delete(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	var user User
	userID := mux.Vars(r)["userID"]

	id, err := helpers.ParseIDs([]string{userID})
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, "failed to delete user")
		return
	}

	user.ID = id[0]
	err = u.service.Delete(&user)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, "failed to delete user")
		return
	}

	helpers.JSONResponse(w, http.StatusNoContent, map[string]string{
		"message": "successfully deleted user",
	})
	return
}

func (u *handler) CreateAdmin(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	_, err := u.service.Create(&user)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, map[string]string{
		"message": "successfully created admin",
		"Email": user.Email,
		"Name": user.Name,
	})

	return
}

func (u *handler) Login(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	err := u.service.Login(&user, user.Password)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	m := map[string]interface{}{
		"id": user.ID,
		"email": user.Email,
	}

	token, err := helpers.CreateToken(m)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, map[string]string{
		"message": "successfully logged in",
		"Token": token,
		"Email": user.Email,
		"Name": user.Name,
	})
	return
}