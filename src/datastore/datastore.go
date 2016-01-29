package datastore

type ServiceData struct {
	Servicestatus bool
}

const SeriveDataQueueSize = 10

type ServiceDataQueue struct {
     Queue  [SeriveDataQueueSize]*ServiceData
     pos    int
}

func (s *ServiceDataQueue ) PushBack(val *ServiceData) {
     if s.pos == SeriveDataQueueSize {
	s.pos = 0
     }
     s.Queue[s.pos] = val
}

func (s *ServiceDataQueue ) Front() (*ServiceData) {
     return s.Queue[s.pos]
}

type DataStore struct {
  ServiceMap map[string]*ServiceDataQueue
}

func (d *DataStore) Init() {
        d.ServiceMap = make(map[string]*ServiceDataQueue)
}

func (d *DataStore) Update(service string, sdata *ServiceData) error {
        if sdq, ok := d.ServiceMap[service]; ok {
                sdq.PushBack(sdata)
        } else {
                tmp := ServiceDataQueue{pos:0}
		tmp.PushBack(sdata)
                d.ServiceMap[service] = &tmp
        }
	return nil
}
