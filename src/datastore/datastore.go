package datastore

type ServiceData struct {
	Servicestatus bool
}

var DataStore map[string]*[]*ServiceData
