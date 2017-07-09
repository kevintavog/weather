package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-playground/lars"
	"github.com/ian-kent/go-log/log"
	"github.com/laurent22/ical-go"
)

var teamsOverride = map[string]string{
	"Chicago Fire":           "Chicago",
	"Colorado Rapids":        "Rapids",
	"Columbus Crew":          "Crew",
	"D.C. United":            "DC",
	"Eintracht Frankfurt":    "Frankfurt",
	"FC Dallas":              "Dallas",
	"Houston Dynamo":         "Houston",
	"LA Galaxy":              "Galaxy",
	"Minnesota United FC":    "Loons",
	"New York City FC":       "NYCFC",
	"Orlando City SC":        "Orlando",
	"Philadelphia Union":     "Philly",
	"Portland Timbers":       "Timbers",
	"Real Salt Lake":         "RSL",
	"San Jose Earthquakes":   "San Jose",
	"Seattle Sounders FC":    "Sounders",
	"Sporting Kansas City":   "SKC",
	"Vancouver Whitecaps FC": "Whitecaps",
}

func convertTeamName(teamName string) string {
	override, ok := teamsOverride[teamName]
	if ok {
		return override
	}

	log.Warn("missing conversion for '%s'", teamName)
	return teamName
}

func soccerScheduel(c lars.Context) {
	wc := c.(*WeatherContext)

	wc.FieldLogger.Time("soccerSchedule", func() {

		response := make(map[string]interface{})
		fileInfo, err := getSoundersScheduleFile()
		propogateError(err, fmt.Sprintf("Unable to read from %s", soundersScheduleDirectory))

		if fileInfo != nil {
			fileData, err := ioutil.ReadFile(filepath.Join(soundersScheduleDirectory, fileInfo.Name()))
			propogateError(err, fileInfo.Name())

			calendar, err := ical.ParseCalendar(string(fileData))
			propogateError(err, "Unable to parse calender")
			games := make([]map[string]interface{}, 0)
			for _, g := range calendar.ChildrenByName("VEVENT") {

				var defTime time.Time

				summary := g.PropString("SUMMARY", "")
				tokens := strings.Split(summary, "vs.")
				if len(tokens) != 2 {
					log.Error("Unable to parse '%s'", summary)
					continue
				}

				startTime := g.PropDate("DTSTART", defTime)

				gmap := make(map[string]interface{})
				gmap["startTime"] = startTime.Unix()
				gmap["homeTeam"] = convertTeamName(strings.Trim(tokens[0], " "))
				gmap["awayTeam"] = convertTeamName(strings.Trim(tokens[1], " "))

				games = append(games, gmap)
			}

			response["sounders"] = games
		}

		reignGames := make([]map[string]interface{}, 0)
		for _, g := range getReignSchedule() {
			gmap := make(map[string]interface{})
			gmap["startTime"] = g.StartTime.Unix()
			gmap["homeTeam"] = g.HomeTeam
			gmap["awayTeam"] = g.AwayTeam
			reignGames = append(reignGames, gmap)
		}

		response["reign"] = reignGames

		wc.WriteResponse(response)
	})
}
