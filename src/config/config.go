package config

type ServiceConfig struct {
        Servicename string             `json:"servicename"`
        Algorithm                      `json:"algorithm"`
        Servers []Server               `json:"servers"`
}

type Server struct {
        Host    string     `toml:"host"`
        Uri     string     `toml:"uri"`
}

type Algorithm struct {            
	Name string               `json:"name"`
        Ttl  int                  `json:"ttl"`
}
