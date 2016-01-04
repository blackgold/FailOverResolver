package main

import (
	"config"
	"fmt"
)

func main() {
	cnf, err := config.Parse("config/config.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cnf)
}
