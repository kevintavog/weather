package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/ian-kent/go-log/log"
	"github.com/laurent22/ical-go"
	"gopkg.in/resty.v0"
)

// Responsible for collecting the Sounders schedule once a day
func soundersScheduleCollector() {

	fetchSoundersSchedule()

	c := time.Tick(24 * time.Hour)
	for _ = range c {
		fetchSoundersSchedule()
	}
}

func fetchSoundersSchedule() {
	resp, err := resty.R().Get("http://ics.ecal.com/ecal-sub/59276d0273a5ca7e0e8b4567/Seattle%20Sounders.ics")
	if err != nil {
		log.Error("Sounders schedule request failed with an error: %s", err)
		return
	}

	if resp.StatusCode() != http.StatusOK {
		log.Error("Sounders Schedule request failed with status: %d", resp.StatusCode())
		log.Error("%s", resp.Body())
		return
	}

	// Make sure the file is parseable before saving it...
	_, err = ical.ParseCalendar(resp.String())
	if err != nil {
		log.Error("Failed parsing Sounders Schedule: '%s'", resp.String())
		return
	}

	// Save as the current file - so we can fall back to it if there are problems contacting
	// the server
	filename := filepath.Join(soundersScheduleDirectory, "sounders-schedule")

	log.Info("Saving file %s", filename)
	err = ioutil.WriteFile(filename, resp.Body(), os.ModePerm)
	if err != nil {
		log.Error("Failed writing to file %s: %s", filename, err)
		return
	}
}

func getSoundersScheduleFile() (os.FileInfo, error) {
	return os.Stat(filepath.Join(soundersScheduleDirectory, "sounders-schedule"))
}
