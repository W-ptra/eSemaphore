package database

import (
	"eSemaphore-backend/model"
)

func (db *Database) CreateUser(user model.User) (model.User,error) {
	if err := db.Connection.Create(&user).Error; err !=nil{
		return model.User{},err
	}
	return user,nil
}

func (db *Database) GetUserByEmail(email string) (model.User,error) {
	var user model.User
	err := db.Connection.Where("email = ?",email).First(&user).Error
	if err != nil{
		return model.User{},nil
	}
	return user,nil
}