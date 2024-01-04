package main

import (
	"encoding/json"
	"io"
	"io/fs"
	"log"
	"mime"
	"net/http"
	"os"
	"slices"
)

const ADDON_LOG_FILE string = "addon_log.txt"

type Config struct {
	Misc []string `json:"misc"`
	Characters []string `json:"characters"`
	Maps []string `json:"maps"`
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

func handleAddons(addonLinks []string, path string, installPath string) error {
	fullPath := installPath + path
	if addonLinks != nil {
		os.Mkdir(fullPath, os.ModePerm)
		for _, link := range addonLinks {
			addonLog, err := getAddonLog(installPath)
			if err != nil {
				return err
			}
			_, exists := addonLog[link]
			if !exists {
				filename, err := downloadAddon(fullPath, link)
				if err != nil {
					log.Printf("Could not download file")
					return err
				}
				addonLog[link] = filename
			}
		}
	}
	return nil
}

func getAddonLog(installPath string) (map[string]string, error) {
	addonLogFile, err := os.ReadFile(installPath + ADDON_LOG_FILE)
	if err != nil {
		return nil, err
	}
	var addonLog map[string]string
	err = json.Unmarshal(addonLogFile, &addonLog)
	if err != nil {
		return nil, err
	}
	return addonLog, nil
}

func writeAddonLog(installPath string, addonLog map[string]string) error {
	addonLogFile, err := json.Marshal(addonLog)
	out, err := os.Create(installPath + ADDON_LOG_FILE)
	io.Copy(out, addonLogFile)
	return nil
}


func downloadAddon(filepath string, link string) (string, error) {
	log.Printf("Downloading: %s", link)
	res, err := http.Get(link)
	if err != nil {
		return "", err
	}
	filename, err := getFileName(res)
	log.Printf("File name: %s", filename)
	defer res.Body.Close()
	out, err := os.Create(filepath + filename)
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, res.Body)
	if err != nil {
		return "", err
	}
	return filename, nil
}

func getFileName(res *http.Response) (string, error) {
	contentDisposition := res.Header.Get("Content-Disposition")
	disposition, params, err := mime.ParseMediaType(contentDisposition)
	if err != nil {
		return "", err
	}
	log.Printf("Media type: %s", disposition)
	filename := params["filename"]
	return filename, nil
}
