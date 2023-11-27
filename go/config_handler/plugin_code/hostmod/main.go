package main

import (
	"encoding/json"
	"errors"
	"log"
)

type configPlugin struct {}

func (c configPlugin) ProcessConfig(configMap map[string]json.RawMessage) (string, error) {
	configJson, ok := configMap["hostmod"]
	if !ok {
		log.Printf("Could not find hostmod config")
		return "", errors.New("Could not find config")
	}
	
	var hostmod = hostmod{}
	err := json.Unmarshal(configJson, &hostmod)
	if err != nil {
		log.Print(err)
		return "", err
	}
	log.Printf("Processing config")
	return mapHostmod(&hostmod), nil
}

var ConfigPlugin configPlugin
