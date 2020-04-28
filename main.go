package main

import (
	"embassy/internal/database"
	"embassy/internal/routers"
	"github.com/codegangsta/negroni"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Main Method
func main() {
	db := database.PostgresConnection()
	router := routers.InitRoutes(db)
	n := negroni.Classic()
	n.UseHandler(router)

	logrus.Info("Server is running")
	err := http.ListenAndServe(":5500", n)
	if err != nil{
		logrus.Fatal(err)
	}
}
