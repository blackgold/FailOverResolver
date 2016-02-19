package main

import (
	"config"
	"datastore"
	"fmt"
	"github.com/gorilla/mux"
	"handlers"
	"net/http"
	"server"
)

func main() {
	var client http.Client
	confArray, err := config.ParseDir("config")
	if err != nil {
		fmt.Println(err)
	}
	var dataStore datastore.DataStore
	dataStore.Init()
	for _, conf := range *confArray {
		for i, _ := range conf.Servers {
			go server.Run(conf, i, &client, &dataStore)
			// add some artificial delay to spread requests
		}
	}
	hand := handlers.Handler{ConfigArray: confArray, Data: &dataStore}
	rtr := mux.NewRouter()
	rtr.HandleFunc("/services", hand.ListServices).Methods("GET")
	rtr.HandleFunc("/services/{servicename}", hand.ListService).Methods("GET")
	rtr.HandleFunc("/services/{servicename}/stats", hand.ListServiceStats).Methods("GET")
	http.Handle("/", rtr)
	fmt.Println(http.ListenAndServe(":80", nil))
}
