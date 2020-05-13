package main

import (
	"Embassy/internal/database"
	"Embassy/internal/domain/chatting"
	"Embassy/internal/routers"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/handlers"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

// Main Method
func main() {
	db := database.PostgresConnection()
	router := routers.InitRoutes(db)
	n := negroni.Classic()
	n.UseHandler(router)

	go chatting.HandleMessages()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Origin", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	logrus.Info(fmt.Sprintf("Server is running on PORT  %s", os.Getenv("PORT")))
	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")),
		handlers.CORS(originsOk, headersOk, methodsOk)(n))

	if err != nil{
		logrus.Fatal(err)
	}
}
