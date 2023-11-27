package main

import (
	"encoding/json"
	"io/fs"
	"log"
	"os"
	"plugin"
	"github.com/kerrgavin/srb2kart-go/config"
)

type Config struct {
	ServerConfigs map[string]json.RawMessage `json:"serverconfigs"`
}

func main() {
	jsonFile, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal("could not load config json file")
	}

	var configFile = Config{}
	err = json.Unmarshal([]byte(jsonFile), &configFile)
	if err != nil {
		log.Fatal(err)
	}

	files, err := os.ReadDir("plugins")
	if err != nil {
		log.Fatal("plugins fucking gone")
	}
	fileNames := getFileNames(files)
	for _, file := range fileNames {
		log.Printf("%s", file)
		plug, err := plugin.Open("./plugins/" + file)
		if err != nil {
			log.Fatal("couldn't create plugin")
		}
		processor, err := plug.Lookup("ConfigPlugin")
		if err != nil {
			log.Fatal("not a ConfigPlugin")
		}
		var configPlugin config.ConfigPlugin
		configPlugin, ok := processor.(config.ConfigPlugin)
		if !ok {
			log.Fatal("couldn't load plugin")
		}
		pluginOut, err := configPlugin.ProcessConfig(configFile.ServerConfigs)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Plugin output: %s", pluginOut)
	}
}

func getFileNames(files []fs.DirEntry) []string {
        fileNames := []string{}
        for _, file := range files {
                if !file.IsDir() {
                        fileNames = append(fileNames, file.Name())
                }
        }
        return fileNames
}

