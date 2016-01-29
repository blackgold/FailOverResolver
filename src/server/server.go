package server

import (
	"config"
	"fmt"
	"net/http"
	"time"
)

func Run(conf *config.ServiceConfig, i int, cli *http.Client) {
	for {
		resp, err := cli.Get(conf.Servers[i].Uri)
		if err != nil {
			fmt.Println("Error :", err)
		} else {
		   defer resp.Body.Close()
		   fmt.Println(conf.Algorithm.Name, conf.Servers[i].Host)
		   time.Sleep(time.Duration(conf.Algorithm.Ttl) * time.Second)
		}
	}
}
