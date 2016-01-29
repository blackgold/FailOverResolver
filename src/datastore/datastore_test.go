package datastore_test

import (
	"datastore"
	"testing"
)

func TestUpdate(t *testing.T) {
	tmp := datastore.DataStore{}
        tmp.Init()
	err := tmp.Update("test", &datastore.ServiceData{Servicestatus:true})
	if err != nil {
		t.Error(err.Error())
	}

	if sqd, ok := tmp.ServiceMap["test"]; !ok {
		t.Error("Test key missing")
	} else {
		if !sqd.Front().Servicestatus {
			t.Error("Expected Servicestatus to be true but got it false")
		}
	}
}
