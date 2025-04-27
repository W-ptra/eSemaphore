package dto

import (
	"eSemaphore-backend/model"
	"time"
)

type CreateChatLine struct {
	Content  string `json:"content"`
}

type Chat struct {
	Id               	string 		`json:"id"`
	ReceiverName     	string 		`json:"receiverName"`
	ReceiverImageUrl 	string 		`json:"receiverImageUrl"`
	LatestMessage    	string		`json:"latestMessage"`
	LatestMessageTime   time.Time	`json:"latestMessageTime"`
}

type ChatLine struct {
	Id			string		`json:"id"`
	ChatId		string		`json:"chatId"`
	Content		string		`json:"content"`
	SenderId	string		`json:"senderId"`
	CreatedAt	time.Time	`json:"createdAt"`
}

func ConvertChatModelsToDtos(
	chatModels []model.Chat,
	receiverName,
	receiverImageUrl,
	latestMessage string,
) []Chat {
	var chatDtos []Chat
	for _,chatModel := range chatModels{
		chatDto := Chat{
			Id: chatModel.Id,
			ReceiverName: receiverName,
			ReceiverImageUrl: receiverImageUrl,
			LatestMessage: latestMessage,
			LatestMessageTime: chatModel.CreatedAt,
		}
		chatDtos = append(chatDtos,chatDto)
	}
	return chatDtos
}

func ConvertChatLineModelsToDtos(
	chatId	string,
	chatLineModels []model.ChatLine,
) []ChatLine {
	var chatLineDtos []ChatLine
	for _,chatLineModel := range chatLineModels{
		chatDto := ChatLine{
			Id: chatLineModel.Id,
			ChatId: chatId,
			Content: chatLineModel.Content,
			SenderId: "test",
			CreatedAt: chatLineModel.CreatedAt,
		}
		chatLineDtos = append(chatLineDtos, chatDto)
	}
	return chatLineDtos
}