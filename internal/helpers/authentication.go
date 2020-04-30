package helpers

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
)

func GetUserDataFromMedia(r *http.Request, url string, config *oauth2.Config) ([]byte, error) {
	code := r.FormValue("code")
	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	if !token.Valid() {
		return nil, fmt.Errorf("failed to verify token")
	}

	response, err := http.Get(url+token.AccessToken)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil{
		return nil, err
	}

	return data, nil
}
