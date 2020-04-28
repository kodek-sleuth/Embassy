package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
)


func PostgresConnection() *gorm.DB {
	db, err := gorm.Open("postgres","postgres://postgres:kevina52@localhost:5432/embassy?sslmode=disable")
	if err != nil{
		logrus.Fatal(err)
	}
	logrus.Info("PSQL client connected")
	return db
}