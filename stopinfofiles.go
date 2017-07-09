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

// Filenames match a pattern: stop-YYYY-MM-DD-HH-mm.json (filtered in 'getAllFiles')
var validStopInfoDataFilename = regexp.MustCompile(`^stop-\d{4}-\d{2}-\d{2}-\d{2}-\d{2}\.json$`)

// Responsible for collecting data files from OneBusAway at the appropriate time
// Currently collects files every minute
func stopInfoDataCollector() {

	fetchStopInfo()

	c := time.Tick(1 * time.Minute)
	for _ = range c {
		fetchStopInfo()
	}
}

func fetchStopInfo() {
	url := fmt.Sprintf(
		"http://api.pugetsound.onebusaway.org/api/where/arrivals-and-departures-for-stop/%s.json?key=%s",
		Current.OneBusAwayStop, Current.OneBusAwayKey)
	resp, err := resty.R().Get(url)
	if err != nil {
		log.Error("OneBusAway request failed with an error: %s", err)
		return
	}

	if resp.StatusCode() != http.StatusOK {
		log.Error("OneBusAway request failed with status: %d", resp.StatusCode())
		log.Error("%s", resp.Body())
		return
	}

	// FYI: The timestamp is when we retrieved the file, not when the stop info was updated
	now := time.Now()
	filename := filepath.Join(
		stopInfoDataDirectory,
		fmt.Sprintf("stop-%04d-%02d-%02d-%02d-%02d.json", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute()))

	log.Info("Saving file %s", filename)
	err = ioutil.WriteFile(filename, resp.Body(), os.ModePerm)
	if err != nil {
		log.Error("Failed writing to file %s: %s", filename, err)
		return
	}
}

func pruneStopInfoFiles() {
	files, err := getAllStopInfoFiles()
	if err != nil {
		log.Error("Failed getting latest file: %s", err)
		return
	}

	cutoffTime := time.Now().Add(-20 * time.Minute)
	for _, f := range files {
		if f.ModTime().Before(cutoffTime) {
			log.Info("Deleting expired file: %s", f.Name())
			fullname := filepath.Join(stopInfoDataDirectory, f.Name())
			err = os.Remove(fullname)
			if err != nil {
				log.Error("Failed deleting: %s (%s)", fullname, err)
			}
		}
	}
}

func getAllStopInfoFiles() ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(stopInfoDataDirectory)
	if err != nil {
		return nil, err
	}

	filtered := []os.FileInfo{}
	for _, f := range files {
		if validStopInfoDataFilename.MatchString(f.Name()) {
			filtered = append(filtered, f)
		}
	}

	return filtered, nil
}

func getLatestStopInfo() (os.FileInfo, int, error) {

	var allFiles, err = getAllStopInfoFiles()
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
