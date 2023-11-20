package main

import (
	"log"
)

type configPlugin struct {}

func (c configPlugin) ProcessConfig(configJson string) string {
	log.Printf("Processing config")
	return "this is some example text"
}

var ConfigPlugin configPlugin
