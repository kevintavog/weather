package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"github.com/ian-kent/go-log/log"
	"github.com/spf13/viper"
)

type Configuration struct {
	DarkSkyKey     string
	Latitude       float32
	Longitude      float32
	OneBusAwayKey  string
	OneBusAwayStop string
}

var Current Configuration

func ReadConfiguration() {

	configDirectory := ConfigDirectory

	configFile := path.Join(configDirectory, "rangic.weather")
	_, err := os.Stat(configFile)
	if err != nil {
		defaults := &Configuration{
			DarkSkyKey:     "key goes here",
			Latitude:       0.0,
			Longitude:      0.0,
			OneBusAwayKey:  "key goes here",
			OneBusAwayStop: "stop goes here",
		}
		json, jerr := json.Marshal(defaults)
		if jerr != nil {
			log.Fatal("Config file (%s) doesn't exist; attempt to write defaults failed: %s", configFile, jerr.Error())
		}

		werr := ioutil.WriteFile(configFile, json, os.ModePerm)
		if werr != nil {
			log.Fatal("Config file (%s) doesn't exist; attempt to write defaults failed: %s", configFile, werr.Error())
		} else {
			log.Fatal("Config file (%s) doesn't exist; one was written with defaults", configFile)
		}
	} else {
		viper.SetConfigFile(configFile)
		viper.SetConfigType("json")
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal("Error reading config file (%s): %s", configFile, err.Error())
		}
		err = viper.Unmarshal(&Current)
		if err != nil {
			log.Fatal("Failed converting configuration from (%s): %s", configFile, err.Error())
		}
	}
}
