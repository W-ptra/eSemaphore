package model

import "time"

type ChatLine struct{
	Id			string		`json:"id" gorm:"primaryKey"`
	Content		string		`json:"content"`
	CreatedAt	time.Time	`json:"createdAt"`
	
	UserID		string		`json:"userID"`
	User		User		`gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	ChatID		string		`json:"chatID"`
	Chat		Chat		`gorm:"foreignKey:ChatID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;`
}