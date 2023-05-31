package database

import (
	"go-microservices/common/utilities"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type DBConnection interface {
	Connection() *gorm.DB
}

type service struct{}

func NewDBService() *service{
	return &service{}
}


func (s *service) Connection() *gorm.DB {
	databaseURI := make(chan string, 1)

	databaseURI <- utilities.GodotEnv("DATABASE_URL_DEV")

	db, err := gorm.Open("postgres", <-databaseURI)

	if err != nil {
		defer logrus.Info("Connection to Database Failed")
		logrus.Fatal(err.Error())
	} else {
		logrus.Info("Connection to Database Successfully")
	}

	//  setup db migrations here
	databaseMigrations(db)

	return db
}


func databaseMigrations(db *gorm.DB) {
	//

	logrus.Info("Database migrations")

}

