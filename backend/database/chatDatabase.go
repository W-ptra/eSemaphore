package database

import "eSemaphore-backend/model"

func (db *Database) CreateChats(chat model.Chat) (model.Chat,error) {
	if err := db.Connection.Create(&chat).Error; err != nil{
		return model.Chat{},nil
	}
	return chat,nil
}

func (db *Database) GetChatByUserId(userId string) ([]model.Chat,error) {
	var chats []model.Chat
	err := db.Connection.Where("user_id = ?",userId).Find(&chats).Error; 
	if err != nil{
		return nil,err
	}
	return chats,nil
}