package middlewares

import (
	"Embassy/internal/helpers"
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
	if claims.IsAdmin != true {
		helpers.ErrorResponse(w, http.StatusForbidden,
			"failed to perform action, please contact administration for help")
		return
	}
	next(w, r)
	return
}