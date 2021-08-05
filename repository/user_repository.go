package repository

import (
	"fmt"
	"github.com/airoasis/user/config"
	"github.com/airoasis/user/model/entity"
)

//GetAllUsers Fetch all user data
func GetAllUsers(user *[]entity.User) (err error) {
	if err = config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

//CreateUser ... Insert New data
func CreateUser(user *entity.User) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetUserByID(user *entity.User, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByUsername(user *entity.User, username string) (err error) {
	if err = config.DB.Where("username = ?", username).First(user).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUser ... Update user
func UpdateUser(user *entity.User, id string) (err error) {
	fmt.Println(user)
	config.DB.Save(user)
	return nil
}

//DeleteUser ... Delete user
func DeleteUser(user *entity.User, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(user)
	return nil
}