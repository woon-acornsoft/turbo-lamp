package postgres

import (
	"fmt"

	"github.com/go-pg/pg/v10"
)

type User struct {
	tableName struct{} `pg:"user,alias:user"`
	Id        string
	Name      string
}

func (u User) String() string {
	return fmt.Sprintf("User<%s %s>", u.Id, u.Name)
}

func ExampleDB_Model() {
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "auth",
		Password: "pass",
		Database: "auth",
	})
	defer db.Close()

	// err := createSchema(db)
	// if err != nil {
	// 	panic(err)
	// }

	// Select all users.
	var users []User
	err := db.Model(&users).Select()
	if err != nil {
		panic(err)
	}
	fmt.Println(users)

	user1 := &User{
		Id:   "new",
		Name: "뉴뉴",
	}

	r, er := db.Model(user1).Insert()
	fmt.Println("result : ", r)
	if er != nil {
		panic(er)
	}

	// // Select user by primary key.
	// user := &User{Id: user1.Id}
	// err = db.Model(user).WherePK().Select()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(user)
}

// createSchema creates database schema for User and Story models.
// func createSchema(db *pg.DB) error {
// 	models := []interface{}{
// 		(*User)(nil),
// 	}

// 	for _, model := range models {
// 		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
// 			Temp: true,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
