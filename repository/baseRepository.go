package repository

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/woon-acornsoft/turbo-lamp/config"
)

type repository struct {
	db *pg.DB
	c  *config.Config
}

var db *pg.DB

func Connect(c *config.Config) {

	db = pg.Connect(&pg.Options{
		Addr:     c.Db.Address,
		User:     c.Db.User,
		Password: c.Db.Password,
		Database: c.Db.Database,
	})
}

func Close() {
	if db != nil {
		db.Close()
	}
}

func Check() {
	var num int
	_, err := db.Query(pg.Scan(&num), "SELECT ?", 1)
	if err != nil {
		panic(err)
	}
	fmt.Println("simple:", num)
}
