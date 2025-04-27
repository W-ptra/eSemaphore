package database

import (
	"eSemaphore-backend/model"
)

func (db *Database) CreateUser(user model.User) (model.User,error) {
	if err := db.connection.Create(&user).Error; err !=nil{
		return model.User{},err
	}
	return user,nil
}

func (db *Database) GetUserByEmail(email string) (model.User,error) {
	var user model.User
	err := db.connection.Where("email = ?",email).First(&user).Error
	if err != nil{
		return model.User{},err
	}
	return user,nil
}

func (db *Database) UpdateUserProfile(user model.User) error {
	err := db.connection.Model(&user).Updates(model.User{
		Name: user.Name,
		ImageUrl: user.ImageUrl,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) UpdateUserEmail(user model.User) error {
	err := db.connection.Model(&user).Updates(model.User{
		Email: user.Email,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) UpdateUserPassword(user model.User) error {
	err := db.connection.Model(&user).Updates(model.User{
		Password: user.Password,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) GetUserById(id string) (model.User,error) {
	var user model.User
	err := db.connection.Where("id = ?",id).First(&user).Error
	if err != nil{
		return model.User{},err
	}
	return user,nil
}

