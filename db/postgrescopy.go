package postgres

import (
	"fmt"

	"github.com/woon-acornsoft/turbo-lamp/model"
)

func GetUsers(id string) model.User {
	db.Begin()
	return model.User{}
}

func GetUser(id string) (*model.User, error) {
	fmt.Println("GetUser ==>> start")
	user := model.User{Id: "no"}
	// user.Id = "a"
	fmt.Println("user ==>> ", user)
	// fmt.Println(&user)
	err := db.Model(&user).WherePK().Select()
	fmt.Println(err)
	fmt.Println("user ==>> ", user)
	return &user, err
}
