package services

import (
	"net/http"
	"os"

	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-sms-service/constants"
	"github.com/bitokss/bitok-sms-service/domains/v1"
	repositories "github.com/bitokss/bitok-sms-service/repositories/postgres/v1"
	"github.com/kavenegar/kavenegar-go"
)

var MessagesService messagesServiceIntreface = &messageService{}

type messagesServiceIntreface interface {
	SendMessage(message domains.MessageRequest) (rest_response.RestResp, rest_response.RestResp)
	SendOtp(message domains.OtpRequest) (rest_response.RestResp, rest_response.RestResp)
}

type messageService struct{}

func (*messageService) SendMessage(message domains.MessageRequest) (rest_response.RestResp, rest_response.RestResp) {
	api := kavenegar.New(os.Getenv(constants.KavenegarApiKey))
	if _, err := api.Message.Send("", message.To, message.Message, nil); err != nil {
		return nil, rest_response.NewInternalServerError(err.Error(), nil)
	}
	m := domains.Message{
		Message: message.Message,
		To:      message.To,
	}
	mResp, err := repositories.MessagesRespository.Create(m)
	if err != nil {
		return nil, err
	}
	return rest_response.NewSuccessResponse(constants.MessageSentSuccessfully, mResp, http.StatusOK), nil
}

func (*messageService) SendOtp(message domains.OtpRequest) (rest_response.RestResp, rest_response.RestResp) {
	api := kavenegar.New(os.Getenv(constants.KavenegarApiKey))
	template := os.Getenv(constants.KavenegarOtpTemplate)
	params := &kavenegar.VerifyLookupParam{}
	if _, err := api.Verify.Lookup(message.To, template, message.Code, params); err != nil {
		return nil, rest_response.NewInternalServerError(err.Error(), nil)
	}
	mResp := domains.OtpResp{
		To:      message.To,
		Code:    message.Code,
	}
	return rest_response.NewSuccessResponse(constants.MessageSentSuccessfully, mResp, http.StatusOK), nil
}
