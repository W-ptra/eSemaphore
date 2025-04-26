package model

import "time"

type Chat struct {
	Id        	string     	`json:"id" gorm:"primaryKey"`
	CreatedAt 	time.Time	`json:"createdAt"`
	
	Users		[]User		`json:"users" gorm:"many2many:user_chats"`
	Messages    []ChatLine 	`json:"messages" gorm:"foreignKey:ChatID"`
}