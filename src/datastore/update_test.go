package datastore_test

import (
	"datastore"
	"testing"
)

func TestUpdate(t *testing.T) {
	datastore.Init()
	err := datastore.Update("test", &datastore.ServiceData{true})
	if err != nil {
		t.Error(err.Error())
	}

	if serviceDataPtrArray, ok := datastore.DataStore["test"]; !ok {
		t.Error("Test key missing")
	} else {
		if !(*serviceDataPtrArray)[0].Servicestatus {
			t.Error("Expected Servicestatus to be true but got it false")
		}
	}
}
