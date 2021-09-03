package repositories

import (
	"fmt"
	"log"
	"os"

	"github.com/bitokss/bitok-sms-service/constants"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB = PostgresInit()
)

func PostgresInit() *gorm.DB {
	host, user, password, name, port := os.Getenv(constants.DBHost),
		os.Getenv(constants.DBUser), os.Getenv(constants.DBPassword),
		os.Getenv(constants.DBName), os.Getenv(constants.DBPort)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, name, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
