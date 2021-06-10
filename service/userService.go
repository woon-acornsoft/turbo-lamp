package service

import (
	"fmt"

	"github.com/woon-acornsoft/turbo-lamp/model"
	"github.com/woon-acornsoft/turbo-lamp/repository"
)

func (service) GetUserById() *model.User {
	return nil
}

func Users() *[]model.User {
	users, err := repository.GetUsers()
	if err == nil {
		fmt.Println("users error TT")
		fmt.Println(err)
	}
	return users
}

func User(id string) *model.User {
	user, err := repository.GetUser(id)
	if err == nil {
		fmt.Println(err)
	}
	return user
}
