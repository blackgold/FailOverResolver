package datastore

import (
	"errors"
	"sync"
)

type ServerData struct {
	Serverstatus bool
}

const ServerDataQueueSize = 60 // make this configurable

type ServerDataQueue struct {
	Queue [ServerDataQueueSize]*ServerData
	Pos   int
}

func (s *ServerDataQueue) PushBack(val *ServerData) {
	if s.Pos == ServerDataQueueSize {
		s.Pos = 0
	}
	s.Queue[s.Pos] = val
	s.Pos += 1
}

type ServiceData struct {
	ServiceDataMap map[string]*ServerDataQueue
}

type DataStore struct {
	DataStoreMap map[string]*ServiceData
	mtx          sync.Mutex
}

func (d *DataStore) Init() {
	d.DataStoreMap = make(map[string]*ServiceData)
}

func (d *DataStore) Update(service string, server string, sdata *ServerData) error {
	d.mtx.Lock()
	defer d.mtx.Unlock()
	if sd, ok := d.DataStoreMap[service]; ok {
		if sdq, ok := sd.ServiceDataMap[server]; ok {
			sdq.PushBack(sdata)
		} else {
			tmp := ServerDataQueue{Pos: 0}
			tmp.PushBack(sdata)
			sd.ServiceDataMap[server] = &tmp
		}
	} else {
		sdq := ServerDataQueue{Pos: 0}
		sdq.PushBack(sdata)
		var sd ServiceData
		sd.ServiceDataMap = make(map[string]*ServerDataQueue)
		sd.ServiceDataMap[server] = &sdq
		d.DataStoreMap[service] = &sd
	}
	return nil
}

func (d *DataStore) Get(service string) (*ServiceData, error) {
	d.mtx.Lock()
	defer d.mtx.Unlock()
	if sd, ok := d.DataStoreMap[service]; ok {
		return sd, nil
	}
	return nil, errors.New("Service " + service + " Not Configured")
}
