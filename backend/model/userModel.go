package model

import "time"

type User struct{
	Id			string			`json:"id" gorm:"primaryKey"`
	Name		string			`json:"name"`
	Email		string			`json:"email" gorm:"index"`
	Password	string			`json:"password"`
	ImageUrl	string			`json:"imageUrl"`	
	CreatedAt	time.Time		`json:"createdAt"`

	Chats		[]Chat			`json:"Chats" gorm:"many2many:user_chats"`
}