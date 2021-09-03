package app

import (
	"net/http"
	"os"

	"github.com/bitokss/bitok-sms-service/constants"
	"github.com/bitokss/bitok-sms-service/domains/v1"
	repositories "github.com/bitokss/bitok-sms-service/repositories/postgres/v1"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var (
	e *echo.Echo
)

// Validator is struct which hold validator instance and spread it in whole application
type Validator struct {
	validator *validator.Validate
}

// Validate is a method which specifies how to face with invalid inputs
func (v *Validator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, constants.InvalidInputErr)
	}
	return nil
}

// StartApp is function to Start application
func StartApp(port string) {
	e = echo.New()
	if os.Getenv(constants.KavenegarApiKey) == "" {
		e.Logger.Fatalf("you should set %s enviroment variable", constants.KavenegarApiKey)
	}
	// validate inputs using go-playground package
	e.Validator = &Validator{validator: validator.New()}
	urlMapper()
	// initialize postgres and get db instance
	db := repositories.PostgresInit()
	// autoMigrate will automatically create tables using domains
	err := db.AutoMigrate(&domains.Message{})
	if err != nil {
		e.Logger.Error(err)
	}
	// start echo server
	e.Logger.Error(e.Start(port))
}
