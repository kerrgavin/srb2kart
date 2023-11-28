package main

import (
	"encoding/json"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"slices"
)

type Addon struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

type Config struct {
	Misc []Addon `json:"misc"`
	Characters []Addon `json:"characters"`
	Maps []Addon `json:"maps"`
}

func main() {
	configFileName := os.Args[1]
	installPath := os.Args[2]

	if configFileName == "" || installPath == "" {
		log.Fatal("Missing arguments")
	}
	configFile, err := os.ReadFile(configFileName)
	if err != nil {
		log.Fatal("Could not open config file")
	}

	var config = Config{}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatal("Could not parse config file")
	}

	err = handleAddons(config.Misc, "misc/", installPath)
	if err != nil {
		log.Fatal("Could not process misc addons")
	}
	err = handleAddons(config.Characters, "characters/", installPath)
	if err != nil {
		log.Fatal("Could not process characters addons")
	}
	err = handleAddons(config.Maps, "maps/", installPath)
	if err != nil {
		log.Fatal("Could not process maps addons")
	}
}

func handleAddons(addons []Addon, path string, installPath string) error {
	fullPath := installPath + path
	if addons != nil {
		os.Mkdir(fullPath, os.ModePerm)
		for _, addon := range addons {
			files, err := os.ReadDir(fullPath)
			if err != nil {
				log.Println("Could not read directory")
				return err
			}
			fileNames := getFileNames(files)
			if !slices.Contains[[]string, string](fileNames, addon.Name) {
				err = downloadAddon(fullPath, addon)
				if err != nil {
					log.Println("Could not download file")
					return err
				}
			}
		}
	}
	return nil
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

func downloadAddon(filepath string, addon Addon) error {
	log.Printf("Downloading %s from %s", addon.Name, addon.Url)
	res, err := http.Get(addon.Url)
	if err != nil {
		return err
	}
	log.Printf("File name: %s", res.Header.Get("Content-Disposition"))
	defer res.Body.Close()
	out, err := os.Create(filepath + addon.Name)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, res.Body)
	return err
}
