package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	ZincServer   string `json:"zinc_server"`
	ZincUser     string `json:"zinc_user"`
	ZincPassword string `json:"zinc_password"`
	ZincIndex    string `json:"zinc_index"`
}

func NewConfig(configFileName string) (*Config, error) {
	datos, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return nil, err
	}
	var c *Config
	err = json.Unmarshal(datos, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
