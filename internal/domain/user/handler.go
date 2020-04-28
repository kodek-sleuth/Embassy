package user

import (
	"context"
	"crypto/rand"
	//"embassy/internal/helpers"
	"encoding/base64"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// Methods to be consumed by handler
type Handler interface {
	Create(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	GoogleLogin(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	GoogleCallBack(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
}


type handler struct {
	service Service
}


func NewHandler(service Service) Handler {
	return &handler{
		service,
	}
}

var googleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:5500/api/auth/google/callback",
	ClientID:     os.Getenv("CLIENT_ID"),
	ClientSecret: os.Getenv("CLIENT_SECRET"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

func getUserDataFromGoogle(code string) ([]byte, error) {
	// Use code to get token and get user info from Google.

	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}
	return contents, nil
}

func (u *handler) GoogleLogin(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	// Create oauthState cookie
	oauthState := generateStateOauthCookie(w)
	/*
		AuthCodeURL receive state that is a token to protect the user from CSRF attacks.
		You must always provide a non-empty string and
		validate that it matches the the state query parameter on your redirect callback.
	*/
	y := googleOauthConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, y, http.StatusTemporaryRedirect)
}

func (u *handler) GoogleCallBack(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	// Read oauthState from Cookie
	oauthState, _ := r.Cookie("oauthstate")

	if r.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth google state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	data, err := getUserDataFromGoogle(r.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// GetOrCreate User in your db.
	// Redirect or response with a token.
	// More code .....
	fmt.Fprintf(w, "UserInfo: %s\n", data)
}






























func (u *handler) Create(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	//var user User
	//
	//if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
	//	helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
	//	return
	//}
	//
	//// FetchIP Address
	//forwarded := r.Header.Get("X-FORWARDED-FOR")
	//if forwarded != "" {
	//	helpers.ErrorResponse(w, http.StatusInternalServerError, forwarded)
	//	return
	//}
	//
	//user.IP = r.RemoteAddr
	//result, err := u.service.Create(&user)
	//if err != nil {
	//	helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
	//	return
	//}
	//
	//createdUser := CreatedUser{
	//	Email:result.Email,
	//	Password: result.Password,
	//	IP: result.Password,
	//}

	//helpers.JSONResponse(w, http.StatusCreated, createdUser)
	return
}


//func (u *handler) Login(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
//	var user User
//
//	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
//		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	// FetchIP Address
//	forwarded := r.Header.Get("X-FORWARDED-FOR")
//	if forwarded != "" {
//		helpers.ErrorResponse(w, http.StatusInternalServerError, forwarded)
//		return
//	}
//
//	user.IP = r.RemoteAddr
//	err := u.service.Login(&user, user.Password)
//	if err != nil {
//		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	m := map[string]interface{}{
//		"id": user.ID,
//		"email": user.Email,
//	}
//
//	token, err := helpers.CreateToken(m)
//	if err != nil{
//		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	createdUser := CreatedUser{
//		Time: time.Now(),
//		IP: user.IP,
//		Token: token,
//	}
//
//	helpers.JSONResponse(w, http.StatusCreated, createdUser)
//	return
//}