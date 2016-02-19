package handlers

import (
	"config"
	"datastore"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"resolver"
)

type Services struct {
	Services []string `json:"services"`
}

type Servers struct {
	Hostnames []string `json:"hostnames"`
}

type Stat struct {
	Servers []Server `json:"servers"`
}

type Server struct {
	Hostname string `json:"hostname"`
	Status   []bool `json:"Status"`
}

type Handler struct {
	ConfigArray *[]*config.ServiceConfig
	Data        *datastore.DataStore
}

func (h *Handler) ListServices(w http.ResponseWriter, r *http.Request) {

	var res Services
	for _, conf := range *h.ConfigArray {
		res.Services = append(res.Services, conf.Servicename)
	}

	out, err := json.Marshal(res)
	if err != nil {
		log.Println("ListServices: json marshall failed  error: ", err)
		http.Error(w, "{\"Error\":\"Internal Server Error\"}", 500)
	} else {
		fmt.Fprintf(w, "%s", string(out))
	}
}

func (h *Handler) ListService(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	for _, conf := range *h.ConfigArray {
		if conf.Servicename == vars["servicename"] {
			out, err := json.Marshal(conf)
			if err != nil {
				log.Println("ListService: json marshall failed  error: ", err)
				http.Error(w, "{\"Error\":\"Internal Server Error\"}", 500)
				return
			} else {
				fmt.Fprintf(w, "%s", string(out))
			}
			return
		}
	}
	log.Println("ListService: service not found in config", vars["servicename"])
	http.Error(w, "{\"Error\":\"Internal Server Error\"}", 500)
}

func (h *Handler) ListServiceStats(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	for _, conf := range *h.ConfigArray {
		if conf.Servicename == vars["servicename"] {
			sd, err := h.Data.Get(conf.Servicename)
			if sd != nil && err == nil {
				var res Stat
				for key, val := range sd.ServiceDataMap {
					var ser Server
					ser.Hostname = key
					for i := 0; i < val.Pos; i++ {
						ser.Status = append(ser.Status, val.Queue[i].Serverstatus)
					}
					res.Servers = append(res.Servers, ser)
				}
				out, err := json.Marshal(res)
				if err != nil {
					log.Println("ListServiceStats: json marshall failed  error: ", err)
					http.Error(w, "{\"Error\":\"Internal Server Error\"}", 500)
					return
				} else {
					fmt.Fprintf(w, "%s", string(out))
					return
				}
			} else {
				fmt.Println("serive missing in datastore")
			}
		}
	}
	log.Println("ListServiceStats: service not found in config", vars["servicename"])
	http.Error(w, "{\"Error\":\"Internal Server Error\"}", 500)
}

func (h *Handler) Resolve(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rslvr := resolver.Resolver{ConfigArray: h.ConfigArray, Data: h.Data}
	lst, err := rslvr.Resolve(vars["servicename"])
	if err == nil {
		var res Servers
		res.Hostnames = lst
		out, err := json.Marshal(res)
		if err != nil {
			log.Println("Resolve: json marshall failed  error: ", err)
			http.Error(w, "{\"Error\":\"Internal Server Error\"}", 500)
			return
		} else {
			fmt.Fprintf(w, "%s", string(out))
			return
		}
	}
	log.Println("Resolve error", err)
	http.Error(w, "{\"Error\":\"Internal Server Error\"}", 500)
}
