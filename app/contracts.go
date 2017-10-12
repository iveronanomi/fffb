package app

import "github.com/ereminIvan/fffb/model"

type IFBService interface {
	LatestMessages() []model.Message
	ReadMessages() []string
}

type ITelegramService interface {
	SendMessage(message model.Message)
}