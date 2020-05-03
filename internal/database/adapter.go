package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"os"
)


func PostgresConnection() *gorm.DB {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil{
		logrus.Fatal(err)
	}
	logrus.Info("Postgresql client connected")
	return db
}