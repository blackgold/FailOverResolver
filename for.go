package main

import (
	"config"
	"fmt"
)

func main() {
	confArray, err := config.ParseDir("config")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("%v",confArray)
}
