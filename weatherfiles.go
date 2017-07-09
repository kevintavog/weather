package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/ian-kent/go-log/log"
	"gopkg.in/resty.v0"
)

// Filenames match a pattern: weather-YYYY-MM-DD-HH-mm.json (filtered in 'getAllFiles')
var validWeatherDataFilename = regexp.MustCompile(`^weather-\d{4}-\d{2}-\d{2}-\d{2}-\d{2}\.json$`)

// Responsible for collecting data files from DarkSky at the appropriate time
// Currently collects files every 30 minutes, plus at startup if needed
func weatherDataCollector() {

	thirtyMinuteWeatherCheck()

	c := time.Tick(30 * time.Second)
	for _ = range c {
		thirtyMinuteWeatherCheck()
	}
}

// fetch data every 30 minutes, regardless of how many times this method is called.
// Data will be fetched if:
//		1. It's 0 or 30 minutes past the hour
//		2. If the 30 minute block is missing the data file
//			a. 0-29 fulfills the 0 block
//			b. 30-59 fulfills the 30 block
//
// Due to #2, when this is called at startup, if the most recent block doesn't have data, then
// data is fetched.
func thirtyMinuteWeatherCheck() {

	fileInfo, _, err := getLatestWeather()
	if err != nil {
		log.Error("Failed getting latest file: %s", err)
		return
	}
	if fileInfo == nil {
		fetchWeatherData()
		return
	}

	now := time.Now()
	nowMinute := now.Minute()
	blockMinute := 0
	if nowMinute >= 0 && nowMinute <= 29 {
		blockMinute = 0
	} else {
		blockMinute = 30
	}
	blockStart := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), blockMinute, 0, 0, now.Location())

	if fileInfo.ModTime().Before(blockStart) {
		fetchWeatherData()
	}
}

func fetchWeatherData() {
	url := fmt.Sprintf(
		"https://api.darksky.net/forecast/%s/%f,%f?exclude=minutely,alerts,flags&extend=hourly",
		Current.DarkSkyKey, Current.Latitude, Current.Longitude)
	resp, err := resty.R().Get(url)
	if err != nil {
		log.Error("DarkSky request failed with an error: %s", err)
		return
	}

	if resp.StatusCode() != http.StatusOK {
		log.Error("DarkSky request failed with status: %d", resp.StatusCode())
		log.Error("%s", resp.Body())
		return
	}

	// FYI: The timestamp is when we retrieved the file, not when the forecast was generated
	now := time.Now()
	filename := filepath.Join(
		weatherDataDirectory,
		fmt.Sprintf("weather-%04d-%02d-%02d-%02d-%02d.json", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute()))

	log.Info("Saving file %s", filename)
	err = ioutil.WriteFile(filename, resp.Body(), os.ModePerm)
	if err != nil {
		log.Error("Failed writing to file %s: %s", filename, err)
		return
	}
}

func pruneWeatherFiles() {
	files, err := getAllWeatherFiles()
	if err != nil {
		log.Error("Failed getting latest file: %s", err)
		return
	}

	cutoffTime := time.Now().Add(-2 * 24 * time.Hour)
	for _, f := range files {
		if f.ModTime().Before(cutoffTime) {
			log.Info("Deleting expired file: %s", f.Name())
			fullname := filepath.Join(weatherDataDirectory, f.Name())
			err = os.Remove(fullname)
			if err != nil {
				log.Error("Failed deleting: %s (%s)", fullname, err)
			}
		}
	}
}

func getAllWeatherFiles() ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(weatherDataDirectory)
	if err != nil {
		return nil, err
	}

	filtered := []os.FileInfo{}
	for _, f := range files {
		if validWeatherDataFilename.MatchString(f.Name()) {
			filtered = append(filtered, f)
		}
	}

	return filtered, nil
}

func getLatestWeather() (os.FileInfo, int, error) {

	var allFiles, err = getAllWeatherFiles()
	if err != nil {
		return nil, -1, err
	}

	var latest os.FileInfo
	for _, f := range allFiles {
		if latest == nil || f.ModTime().After(latest.ModTime()) {
			latest = f
		}
	}

	return latest, len(allFiles), nil
}
