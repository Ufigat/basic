package main

import (
	service "basic/src/service"
	"fmt"
)

func main() {
	Ivan := service.NewHuman("Ivan")
	fmt.Println(Ivan.GetHumanName())
	fmt.Println("hi all")
}
