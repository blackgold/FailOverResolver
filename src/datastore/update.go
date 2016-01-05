package datastore

func Init() {
	DataStore = make(map[string]*[]*ServiceData)
}

func Update(service string, sdata *ServiceData) error {
	if serviceDataPtrArray, ok := DataStore[service]; ok {
		*serviceDataPtrArray = append(*serviceDataPtrArray, sdata)
	} else {
		var tmp []*ServiceData
		tmp = append(tmp, &ServiceData{true})
		DataStore[service] = &tmp
	}
	return nil
}
