package main

import (
	"config"
	"server"
	"fmt"
        "net/http"
)

func main() {
        var client http.Client
	confArray, err := config.ParseDir("config")
	if err != nil {
		fmt.Println(err)
	}
	for _, conf := range *confArray {
		for i, _ := range conf.Servers {
			go server.Run(conf, i, &client)
		}
	}
	select {}
}
