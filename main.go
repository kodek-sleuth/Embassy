package main

import (
	"Embassy/internal/database"
	"Embassy/internal/domain/chatting"
	"Embassy/internal/routers"
	"fmt"
	"github.com/codegangsta/negroni"
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

	logrus.Info("Server is running")
	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), n)
	if err != nil{
		logrus.Fatal(err)
	}
}
