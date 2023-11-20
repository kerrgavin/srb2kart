package main

import (
	"log"
	"io/fs"
	"os"
	"plugin"
)

type ConfigPlugin interface {
	ProcessConfig(configJson string) string
}

func main() {
	jsonFile, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal("could not load config json file")
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
		pluginOut := configPlugin.ProcessConfig(string(jsonFile))
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

