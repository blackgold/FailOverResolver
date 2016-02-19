package handlers

import (
	"config"
	"datastore"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Services struct {
	Services []string `json:"services"`
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
		log.Println("HandleUser: json marshall failed  error: ", err)
		fmt.Printf("{\"Error\":\"Internal Server Error\"}")
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
				log.Println("HandleUser: json marshall failed  error: ", err)
				fmt.Printf("{\"Error\":\"Internal Server Error\"}")
			} else {
				fmt.Fprintf(w, "%s", string(out))
			}
			return
		}
	}
	fmt.Fprintf(w, "%s", "{Error: no service}")
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
					log.Println("HandleUser: json marshall failed  error: ", err)
					fmt.Printf("{\"Error\":\"Internal Server Error\"}")
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
	fmt.Fprintf(w, "%s", "{Error: no service}")
}
