package server

import (
	"config"
	"datastore"
	"fmt"
	"net/http"
	"time"
)

func Run(conf *config.ServiceConfig, i int, cli *http.Client, dataStore *datastore.DataStore) {
	for {
		resp, err := cli.Get(conf.Servers[i].Uri)
		var ds datastore.ServerData
		if err != nil {
			fmt.Println("Error :", err)
			ds.Serverstatus = false

		} else {
			defer resp.Body.Close()
			if resp.StatusCode != 200 {
				fmt.Println("Error : response code for " + conf.Servicename + " " + resp.Status)
				ds.Serverstatus = false
			} else {
				ds.Serverstatus = true
			}
		}
		dataStore.Update(conf.Servicename, conf.Servers[i].Host, &ds)
		time.Sleep(time.Duration(conf.Algorithm.Ttl) * time.Second)
	}
}
