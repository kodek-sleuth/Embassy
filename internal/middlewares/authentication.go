package middlewares

import (
	"Embassy/internal/helpers"
	"github.com/sirupsen/logrus"
	"net/http"
)

func RequireTokenAuthentication(w http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	_, err := helpers.VerifyToken(r)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusUnauthorized, "unauthorised to perform this action," +
			" please signup/login")
		return
	}
	next(w, r)
	return
}

func RequireAdminRights(w http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	claims, _ := helpers.VerifyToken(r)
	logrus.Println(claims.IsAdmin)
	if claims.IsAdmin != true{
		helpers.ErrorResponse(w, http.StatusForbidden,
			"failed to perform action, please contact administration for help")
		return
	}
	next(w, r)
	return
}

func RequireOwnerRights(w http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	claims, _ := helpers.VerifyToken(r)
	if claims.IsRestaurantOwner != true {
		helpers.ErrorResponse(w, http.StatusForbidden,
			"failed to perform action, please contact administration for help")
		return
	}
	next(w, r)
	return
}