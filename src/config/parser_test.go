package config_test

import (
	"config"
	"testing"
)

func TestParse(t *testing.T) {
	conf, err := config.Parse("../../config/config.json")
	if err != nil {
		t.Error("Expected nil and got error " + err.Error())
	}
	if conf.Servicename != "file" {
		t.Error("Expected file and got " + conf.Servicename)
	}
	if conf.Algorithm.Name != "randomized" {
		t.Error("Expected randomized and got " + conf.Algorithm.Name)
	}
}

func TestParseDir(t *testing.T) {
	confArray, err := config.ParseDir("../../config")
	if err != nil {
		t.Error("Expected nil and got error " + err.Error())
	}
	for _, conf := range *confArray {
		if conf.Servicename != "file" {
			t.Error("Expected file and got " + conf.Servicename)
		}
		if conf.Algorithm.Name != "randomized" {
			t.Error("Expected randomized and got " + conf.Algorithm.Name)
		}
	}
}
