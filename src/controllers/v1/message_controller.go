package controllers

import (
	"os"

	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-sms-service/constants"
	"github.com/bitokss/bitok-sms-service/domains/v1"
	"github.com/bitokss/bitok-sms-service/services/v1"
	"github.com/bitokss/bitok-sms-service/utils"
	"github.com/labstack/echo/v4"
)

var MessagesController messagesControllerIntreface = &messageController{}

type messagesControllerIntreface interface {
	SendMessage(c echo.Context) error
	SendOtp(c echo.Context) error
}

type messageController struct{}

func (*messageController) SendMessage(c echo.Context) error {
	m := new(domains.MessageRequest)
	if err := utils.ValidateAndBind(c, m); err != nil {
		return c.JSON(err.Status(), err)
	}
	if m.Token != os.Getenv(constants.KavenegarApiKey) {
		newErr := rest_response.NewUnauthorizedError(constants.UnAuthorizedErr, nil)
		return c.JSON(newErr.Status(), newErr)
	}
	mResp, err := services.MessagesService.SendMessage(*m)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(mResp.Status(), mResp)
}

func (*messageController) SendOtp(c echo.Context) error {
	m := new(domains.OtpRequest)
	if err := utils.ValidateAndBind(c, m); err != nil {
		return c.JSON(err.Status(), err)
	}
	if m.Token != os.Getenv(constants.KavenegarApiKey) {
		newErr := rest_response.NewUnauthorizedError(constants.UnAuthorizedErr, nil)
		return c.JSON(newErr.Status(), newErr)
	}
	mResp, err := services.MessagesService.SendOtp(*m)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(mResp.Status(), mResp)
}
