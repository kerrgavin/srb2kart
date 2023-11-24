package main

import (
	"encoding/json"
	"io/fs"
	"log"
	"os"
	"plugin"
)

type ConfigPlugin interface {
	ProcessConfig(configMap map[string]json.RawMessage) (string, error)
}

type Config struct {
	ServerConfigs map[string]json.RawMessage `json:"serverconfigs"`
}

func main() {
	jsonFile, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal("could not load config json file")
	}

	var config = Config{}
	err = json.Unmarshal([]byte(jsonFile), &config)
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
		var configPlugin ConfigPlugin
		configPlugin, ok := processor.(ConfigPlugin)
		if !ok {
			log.Fatal("couldn't load plugin")
		}
		pluginOut, err := configPlugin.ProcessConfig(config.ServerConfigs)
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

