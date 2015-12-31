package main

import (
	"fmt"
	"config"
)

func main() {
        cnf, err := config.Parse("config/config.json")
        if err != nil {
                fmt.Println(err)
        }
        fmt.Println(cnf)
}
