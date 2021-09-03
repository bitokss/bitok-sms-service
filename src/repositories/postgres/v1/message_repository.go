package repositories

import (
	"strings"
	"time"

	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-sms-service/constants"
	"github.com/bitokss/bitok-sms-service/domains/v1"
)

var MessagesRespository messagesRespositoryIntreface = &messageRespository{}

type messagesRespositoryIntreface interface {
	Create(message domains.Message) (domains.MessageResp, rest_response.RestResp)
}

type messageRespository struct{}

func (*messageRespository) Create(message domains.Message) (domains.MessageResp, rest_response.RestResp) {
	to := strings.Join(message.To, ",")
	to = "{" + to + "}"
	if err := DB.Exec(`INSERT INTO "messages"("created_at","updated_at","deleted_at","message","to") VALUES (?,?,NULL,?,?)`, time.Now(), time.Now(), message.Message, to).Error; err != nil {
		return domains.MessageResp{}, rest_response.NewInternalServerError(constants.InternalServerErr, nil)
	}
	return serializeMessage(message), nil
}

func serializeMessage(message domains.Message) domains.MessageResp {
	return domains.MessageResp{
		Message: message.Message,
		To:      message.To,
	}
}
