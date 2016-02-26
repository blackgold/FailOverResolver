package server

import (
	"config"
	"datastore"
	"github.com/gorilla/mux"
	"handlers"
	"log"
	"net/http"
	"time"
)

func check(conf *config.ServiceConfig, i int, cli *http.Client, dataStore *datastore.DataStore) {
	for {
		start := time.Now()
		resp, err := cli.Get(conf.Servers[i].Uri)
		if resp != nil {
			defer resp.Body.Close()
		}
		duration := time.Since(start)
		var ds datastore.ServerData
		if err != nil {
			log.Println("Error :", err)
			ds.Serverstatus = false
		} else {
			if resp.StatusCode != 200 {
				log.Println("Error : response code for " + conf.Servicename + " " + resp.Status)
				ds.Serverstatus = false
			} else {
				ds.Serverstatus = true
			}
			ds.DurationInNs = duration.Nanoseconds()
		}
		dataStore.Update(conf.Servicename, conf.Servers[i].Host, &ds)
		time.Sleep(time.Duration(conf.Algorithm.Ttl) * time.Second)
	}
}

func Run(configObj *config.Config, dataStore *datastore.DataStore) {
	var client http.Client
	for _, conf := range configObj.ConfArray {
		for i, _ := range conf.Servers {
			go check(conf, i, &client, dataStore)
			// add some artificial delay to spread requests
		}
	}
	hand := handlers.Handler{ConfigObj: configObj, Data: dataStore}
	rtr := mux.NewRouter()
	rtr.HandleFunc("/services", hand.ListServices).Methods("GET")
	rtr.HandleFunc("/services/{servicename}", hand.ListService).Methods("GET")
	rtr.HandleFunc("/services/{servicename}/stats", hand.ListServiceStats).Methods("GET")
	rtr.HandleFunc("/services/{servicename}/resolve", hand.Resolve).Methods("GET")
	http.Handle("/", rtr)
	log.Println(http.ListenAndServe(":80", nil))
}
