package main

import (
	"config"
	"datastore"
	"log"
	"server"
)

func main() {
	var conf config.Config
	err := conf.ParseDir("config")
	if err != nil {
		log.Fatal(err)
	}
	var dataStore datastore.DataStore
	dataStore.Init()
	go server.Run(&conf, &dataStore)

	select {}
}
