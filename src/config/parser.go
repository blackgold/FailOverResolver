package config

import (
        "io/ioutil"
        "encoding/json"
)

func Parse(file string) (*ServiceConfig,error) {
  config := &ServiceConfig{}

  data, err := ioutil.ReadFile(file)
  if err != nil {
       return nil, err
  }

  err = json.Unmarshal(data, config)
  return config, err
}
