package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"
)

func Parse(file string) (*ServiceConfig, error) {
	config := &ServiceConfig{}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, config)
	return config, err
}

func ParseDir(dir string) (*[]*ServiceConfig, error) {
	var ServiceConfigArray []*ServiceConfig
	fileInfoArray, err := ioutil.ReadDir(dir)
	if err != nil {
		return &ServiceConfigArray, err
	}
	var filecount int = 0
	for _, fileInfo := range fileInfoArray {
		if !fileInfo.IsDir() && strings.HasSuffix(fileInfo.Name(), ".json") {
			filecount++
		}
	}
	if filecount < 1 {
		return &ServiceConfigArray, errors.New("No config files found")
	}
	for _, fileInfo := range fileInfoArray {
		if !fileInfo.IsDir() && strings.HasSuffix(fileInfo.Name(), ".json") {
			config, err := Parse(dir + "/" + fileInfo.Name())
			if err == nil {
				ServiceConfigArray = append(ServiceConfigArray, config)
			}
		}
	}
	return &ServiceConfigArray, nil
}
