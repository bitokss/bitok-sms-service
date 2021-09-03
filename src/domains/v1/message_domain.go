package domains

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Message string   `gorm:"column:message;not null;type:text"`
	To      []string `gorm:"column:to;not null;type:varchar[]"`
}

type MessageRequest struct {
	Message string   `json:"message" validate:"required"`
	To      []string `json:"to" validate:"required"`
	Token   string   `json:"token" validate:"required"`
}

type MessageResp struct {
	Message string   `json:"message"`
	To      []string `json:"to"`
}

type OtpRequest struct {
	To      string `json:"to" validate:"required"`
	Code    string `json:"code" validate:"required"`
	Token   string `json:"token" validate:"required"`
}

type OtpResp struct {
	To      string `json:"to"`
	Code    string `json:"code"`
}
