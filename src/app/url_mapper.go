package app

import (
	"fmt"

	"github.com/bitokss/bitok-sms-service/constants"
	"github.com/bitokss/bitok-sms-service/controllers/v1"
)

func urlMapper() {
	e.POST(fmt.Sprintf(constants.V1Prefix, "message"), controllers.MessagesController.SendMessage)
	e.POST(fmt.Sprintf(constants.V1Prefix, "otp"), controllers.MessagesController.SendOtp)
}
