package repository

import "github.com/woon-acornsoft/turbo-lamp/model"

func GetUsers() (*[]model.User, error) {
	var userList []model.User
	err := db.Model(&userList).Select()
	return &userList, err
}

func GetUser(id string) (*model.User, error) {
	user := model.User{Id: id}
	err := db.Model(&user).WherePK().Select()
	return &user, err
}
