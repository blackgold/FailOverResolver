package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"
)

type Config struct {
	ConfArray []*ServiceConfig
}

func (c *Config) Parse(file string) (*ServiceConfig, error) {
	config := &ServiceConfig{}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, config)
	return config, err
}

func (c *Config) ParseDir(dir string) error {
	fileInfoArray, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	var filecount int = 0
	for _, fileInfo := range fileInfoArray {
		if !fileInfo.IsDir() && strings.HasSuffix(fileInfo.Name(), ".json") {
			filecount++
		}
	}
	if filecount < 1 {
		return errors.New("No config files found")
	}
	for _, fileInfo := range fileInfoArray {
		if !fileInfo.IsDir() && strings.HasSuffix(fileInfo.Name(), ".json") {
			config, err := c.Parse(dir + "/" + fileInfo.Name())
			if err == nil {
				c.ConfArray = append(c.ConfArray, config)
			}
		}
	}
	return nil
}
