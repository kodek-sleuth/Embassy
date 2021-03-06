package helpers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/objx"
	"net/http"
	"os"
	"strings"
	"time"
)

type Claims struct {
	ID uuid.UUID `json:"id"`
	Email string `json:"email"`
	Name string `json:"name"`
	IsAdmin bool `json:"is_admin"`
	jwt.StandardClaims
}

type UserDetails struct {
	ID uuid.UUID
	Email string
	Name string
}


func CreateToken(payload map[string]interface{}) (string, error) {
	load := objx.New(payload)
	userIdStr := fmt.Sprintf("%v", payload["id"])
	parsedUserID, err := uuid.FromString(userIdStr)
	if err != nil{
		return "", err
	}
	expirationTime := time.Now().Add(60 * time.Hour)
	claims := &Claims{
		ID: parsedUserID,
		Email: load.Get("email").Str(),
		Name: load.Get("name").Str(),
		IsAdmin: load.Get("is_admin").Bool(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func VerifyToken(r *http.Request) (*Claims, error) {
	// Initialize a new instance of `Claims`
	claims := &Claims{}
	tokenString := ExtractToken(r)
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return claims, err
	}

	if !tkn.Valid {
		return claims, err
	}

	return claims, nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}