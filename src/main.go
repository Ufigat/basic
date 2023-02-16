package main

import (
	"fmt"

	service "github.com/Ufigat/basic"
)

func main() {
	Ivan := service.NewHuman("Ivan")
	fmt.Println(Ivan.GetHumanName())
	fmt.Println("hi all")
}
