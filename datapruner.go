package main

import (
	"os"
	"path"
	"time"

	"github.com/ian-kent/go-log/log"
)

var stopInfoDataDirectory string
var soundersScheduleDirectory string
var weatherDataDirectory string

func initDataDirectories() {
	stopInfoDataDirectory = path.Join(DataDirectory, "stopinfo")
	soundersScheduleDirectory = path.Join(DataDirectory, "sounders")
	weatherDataDirectory = path.Join(DataDirectory, "weather")

	err := os.MkdirAll(stopInfoDataDirectory, os.ModePerm)
	if err != nil {
		log.Fatal("Unable to create data directory (%s): %s", stopInfoDataDirectory, err.Error())
	}

	err = os.MkdirAll(soundersScheduleDirectory, os.ModePerm)
	if err != nil {
		log.Fatal("Unable to create data directory (%s): %s", soundersScheduleDirectory, err.Error())
	}

	err = os.MkdirAll(weatherDataDirectory, os.ModePerm)
	if err != nil {
		log.Fatal("Unable to create data directory (%s): %s", weatherDataDirectory, err.Error())
	}

}

// Responsible for deleting no longer needed data files, those older than the last few days
// Currently deletes files more than 2 days old
func dataPruner() {

	initDataDirectories()
	pruneWeatherFiles()
	pruneStopInfoFiles()

	c := time.Tick(1 * time.Hour)
	for _ = range c {
		pruneWeatherFiles()
		pruneStopInfoFiles()
	}
}
