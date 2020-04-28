package helpers

import (
	"encoding/json"
	"net/http"
)

func JSONResponse(res http.ResponseWriter, statusCode int, payload interface{}){
	res.Header().Set("Content-Type", "application/json")

	// marshal payload
	resp, _ := json.Marshal(payload)

	// return statusCode
	res.WriteHeader(statusCode)

	// return response
	res.Write(resp)
}

func ErrorResponse(res http.ResponseWriter, statusCode int, message string){
	JSONResponse(res, statusCode, map[string]string{"error":  message} )
}
