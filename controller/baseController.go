package controller

import "fmt"

type controller struct {
}

func init() {
	fmt.Println("controller init func~~~")
}

func Ctl() {
	fmt.Println("func ctl run")
}

func Gogogo() string {
	fmt.Println("go?go?")
	return "gogogo"
}
