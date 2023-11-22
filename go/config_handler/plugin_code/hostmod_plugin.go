package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type hostmod struct {
	Automate bool `json:"automate"`
	Encore int `json:"encore"`
	Specbomb specbomb `json:"specbomb"`
	NameFilter nameFilter `json:"namefilter"`
	Battle battle `json:"battle"`
	Motd motd `json:"motd"`
	Restat restat `json:"restat"`
	Shout shout `json:"shout"`
	Schedule schedule `json:"schedule"`
	Vote vote `json:"vote"`
	CustomCheck customcheck `json:"customcheck"`
	Scoreboard scoreboard `json:"scoreboard"`
	Veto veto `json:"veto"`
}

type specbomb struct {
	Enabled bool `json:"enabled"`
	AntiSoftlock bool `json:"antisoftlock"`
}

type nameFilter struct {
	Mode string `json:"mode"`
	Names []string `json:"names"`
}

type battle struct {
	Timelimit int `json:"timelimit"`
	Bail bool `json:"bail"`
}

type motd struct {
	Enabled bool `json:"enabled"`
	Nag bool `json:"nag"`
	Background string `json:"background"`
	Name string `json:"name"`
	Contact string `json:"contact"`
	Tagline string `json:"tagline"`
}

type restat struct {
	Enabled bool `json:"enabled"`
	Notify bool `json:"notify"`
}

type shout struct {
	Autoshout bool `json:"autoshout"`
	Name string `json:"name"`
	Color string `json:"color"`
}

type schedule struct {
	Enabled bool `json:"enabled"`
	Jobs []string `json:"jobs"`
}

type vote struct {
	Whitelist []string `json:"whitelist"`
	Timer int `json:"timer"`
	Autopass int `json:"autopass"`
	Allowidc bool `json:"allowidc"`
}

type customcheck struct {
	Enabled bool `json:"enabled"`
	CheckOne bool `json:"checkone"`
	CheckTwo bool `json:"checktwo"`
	CheckThree bool `json:"checkthree"`
}

type scoreboard struct {
	Enabled bool `json:"enabled"`
	Humor bool `json:"humor"`
	Lines []string `json:"lines"`
	ModLines []string `json:"modlines"`
}

type veto struct {
	Enabled bool `json:"enabled"`
	Threshold int `json:"threshold"`
	Hellclosed int `json:"hellclosed"`
	Hellopen int `json:"hellopen"`
}

func (hostmod hostmod) getServerConfig() string {
	var serverConfig string = ""
	serverConfig += "#hostmod configs\n"
	serverConfig += fmt.Sprint("hm_automate %s\n", boolToConfig(hostmod.Automate)) 
	serverConfig += fmt.Sprint("hm_encore %d\n", hostmod.Encore)
	serverConfig += hostmod.Motd.getServerConfig()
	return serverConfig
}

func (motd motd) getServerConfig() string {
	var serverConfig string = ""
	serverConfig += fmt.Sprint("hm_motd %s\n", boolToConfig(motd.Enabled))
	serverConfig += fmt.Sprint("hm_motd_nag %s\n", boolToConfig(motd.Nag))
	serverConfig += fmt.Sprint("hm_motd_bg %s\n", motd.Background)
	serverConfig += fmt.Sprint("hm_motd_name %s\n", motd.Name)
	serverConfig += fmt.Sprint("hm_motd_contact %s\n", motd.Contact)
	serverConfig += fmt.Sprint("hm_motd_tagline %s\n", motd.Tagline)
	return serverConfig
}

type configPlugin struct {}

func (c configPlugin) ProcessConfig(configMap map[string]string) (string, error) {
	configJson, ok := configMap["hostmod"]
	if !ok {
		log.Printf("Could not find hostmod config")
		return "", errors.New("Could not find config")
	}
	
	var hostmod = hostmod{}
	err := json.Unmarshal([]byte(configJson), &hostmod)
	if err != nil {
		log.Print(err)
		return "", err
	}
	log.Printf("Processing config")
	return "this is some example text"
}

var ConfigPlugin configPlugin

func boolToConfig(value bool) string {
	if value {
		return "On"
	}
	return "Off"
}
