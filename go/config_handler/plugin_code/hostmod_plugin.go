package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
)

type hostmod struct {
	Automate *bool `json:"automate"`
	Encore *int `json:"encore"`
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
	Enabled *bool `json:"enabled"`
	Nag *bool `json:"nag"`
	Background *string `json:"background"`
	Name *string `json:"name"`
	Contact *string `json:"contact"`
	Tagline *string `json:"tagline"`
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
	var configHelper configHelper
	configHelper.comment("hostmod configs")
	configHelper.configBool("hm_automate", hostmod.Automate)
	configHelper.configInt("hm_encore", hostmod.Encore)
	//serverConfig += hostmod.Specbomb.getServerConfig()
	//serverConfig += hostmod.NameFilter.getServerConfig()
	//serverConfig += hostmod.Battle.getServerConfig()
	hostmod.Motd.getServerConfig(&configHelper)
	//serverConfig += hostmod.Restat.getServerConfig()
	return configHelper.String()
}

func (specbomb specbomb) getServerConfig() string {
	var serverConfig string = ""
	serverConfig += fmt.Sprintf("hm_specbomb %s\n", boolToConfig(specbomb.Enabled))
	serverConfig += fmt.Sprintf("hm_specbomb_antisoftlock %s\n", boolToConfig(specbomb.AntiSoftlock))
	return serverConfig
}

func (nameFilter nameFilter) getServerConfig() string {
	var serverConfig string = ""
	serverConfig += fmt.Sprintf("hm_namefilter_mode %s\n", nameFilter.Mode)
	for _, name := range nameFilter.Names {
		serverConfig += fmt.Sprintf("hm_namefilter %s\n", name)
	}
	return serverConfig
}

func (battle battle) getServerConfig() string {
	var serverConfig string = ""
	serverConfig += fmt.Sprintf("hm_timelimit %d\n", battle.Timelimit)
	serverConfig += fmt.Sprintf("hm_bail %s\n", boolToConfig(battle.Bail))
	return serverConfig
}

func (motd motd) getServerConfig(configHelper *configHelper) {
	configHelper.configBool("hm_motd", motd.Enabled)
	configHelper.configBool("hm_motd_nag", motd.Nag)
	configHelper.configString("hm_motd_bg", motd.Background)
	configHelper.configString("hm_motd_name", motd.Name)
	configHelper.configString("hm_motd_contact", motd.Contact)
	configHelper.configString("hm_motd_tagline", motd.Tagline)
}

func (restat restat) getServerConfig() string {
	var serverConfig string = ""
	serverConfig += fmt.Sprintf("hm_restat %s\n", boolToConfig(restat.Enabled))
	serverConfig += fmt.Sprintf("hm_restat_notify %s\n", boolToConfig(restat.Notify))
	return serverConfig
}

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
	return hostmod.getServerConfig(), nil
}

var ConfigPlugin configPlugin

func boolToConfig(value bool) string {
	if value {
		return "On"
	}
	return "Off"
}

type configHelper struct {
	strings.Builder
}

func (configHelper configHelper) comment(comment string) {
	configHelper.WriteString(fmt.Sprintf("#%s\n", comment))
}

func (configHelper *configHelper) configString(serverConfig string, jsonConfig *string) error {
	if jsonConfig != nil {
		log.Printf("Config: %s %s", serverConfig, *jsonConfig)
		configHelper.WriteString(fmt.Sprintf("%s %s\n", serverConfig, *jsonConfig))
	}
	return nil
}

func (configHelper *configHelper) configBool(serverConfig string, jsonConfig *bool) error {
	if jsonConfig != nil {
		configHelper.WriteString(fmt.Sprintf("%s %s\n", serverConfig, boolToConfig(*jsonConfig)))
	}
	return nil
}

func (configHelper *configHelper) configInt(serverConfig string, jsonConfig *int) error {
	if jsonConfig != nil {
		configHelper.WriteString(fmt.Sprintf("%s %d\n", serverConfig, *jsonConfig))
	}
	return nil
}
