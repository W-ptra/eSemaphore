package database

import "eSemaphore-backend/model"

func (db *Database) CreateChatLine(chatLine model.ChatLine) (model.ChatLine,error){
	if err := db.Connection.Create(&chatLine).Error; err != nil{
		return model.ChatLine{},nil
	}
	return chatLine,nil
}

func (db *Database) GetChatLinesByChatId(chatId string) ([]model.ChatLine,error){
	var chatLines []model.ChatLine
	err := db.Connection.Where("chat_id = ?",chatId).Find(&chatLines).Error
	if err != nil{
		return nil,err
	}
	return chatLines,nil
} 