package main

import (
	"fmt"

	"github.com/woon-acornsoft/turbo-lamp/config"
	"github.com/woon-acornsoft/turbo-lamp/controller"
	// "github.com/woon-acornsoft/turbo-lamp/repository"
	// "github.com/woon-acornsoft/turbo-lamp/router"
)

// pg "github.com/woon-acornsoft/turbo-lamp/db"

func main() {
	fmt.Println("start main")
	// c := config.LoadConfig()
	c := config.GetConfig()
	fmt.Println("config load", c)
	controller.Ctl()
	s := controller.Gogogo()

	fmt.Println("end main...", s)
	// pg.ExampleDB_Model()
	// pg.Connect()
	// pg.GetUser("no")
	// repository.Connect()
	// router.Router()
	// repository.GetUsers()
	// repository.Close()
	t := NewTurbo()
	a := t.tl.GetTurbo("?")
	fmt.Println("a", a)
}

type Turbolamp Turbo

type Turbo struct {
	name string
	tl   *Turbolamp
}

func (s *Turbolamp) GetTurbo(arg string) string {
	fmt.Println("arg", arg)
	return "return " + arg
}

func NewTurbo() *Turbo {
	t := Turbo{name: "turbo"}
	// t.tl = *Turbo
	return &t
}
