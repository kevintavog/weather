package main

import (
	"time"
)

// I don't know of an online source I can download the ReignFC schedule from.
type SoccerGame struct {
	HomeTeam        string
	AwayTeam        string
	StartTime       time.Time
	stringStartTime string
}

var reignGames = []SoccerGame{
	SoccerGame{
		HomeTeam:  "Houston",
		AwayTeam:  "Reign",
		StartTime: time.Date(2017, 5, 27, 13, 0, 0, 0, time.Local).UTC(), //  MAY 27, 2017 | 1:00 PM PT
	},
	SoccerGame{
		HomeTeam:  "Red Stars",
		AwayTeam:  "Reign",
		StartTime: time.Date(2017, 6, 4, 12, 0, 0, 0, time.Local).UTC(), //  JUNE 4, 2017 | 12:00 PM PT
	},
	SoccerGame{
		HomeTeam:  "Kansas",
		AwayTeam:  "Reign",
		StartTime: time.Date(2017, 6, 17, 13, 0, 0, 0, time.Local).UTC(), //  JUNE 17, 2017 | 1:00 PM PT
	},
	SoccerGame{
		HomeTeam:  "Reign",
		AwayTeam:  "Kansas",
		StartTime: time.Date(2017, 6, 24, 19, 0, 0, 0, time.Local).UTC(), //  JUNE 24, 2017 | 7:00 PM PT
	},
	SoccerGame{
		HomeTeam:  "Reign",
		AwayTeam:  "Red Stars",
		StartTime: time.Date(2017, 6, 28, 19, 30, 0, 0, time.Local).UTC(), //  JUNE 28, 2017 | 7:30 PM PT
	},
	SoccerGame{
		HomeTeam:  "Reign",
		AwayTeam:  "Thorns",
		StartTime: time.Date(2017, 7, 1, 19, 0, 0, 0, time.Local).UTC(), //  SAT | JULY 1, 2017 | 7:00 PM PT
	},
	SoccerGame{
		HomeTeam:  "Courage",
		AwayTeam:  "Reign",
		StartTime: time.Date(2017, 7, 8, 13, 0, 0, 0, time.Local).UTC(), //  SAT | JULY 8, 2017 | 1:00 PM PT
	},
	SoccerGame{
		HomeTeam:  "Reign",
		AwayTeam:  "Boston",
		StartTime: time.Date(2017, 7, 15, 19, 0, 0, 0, time.Local).UTC(), //  SAT | JULY 15, 2017 | 7:00 PM PT
	},
	SoccerGame{
		HomeTeam:  "Reign",
		AwayTeam:  "Sky Blue",
		StartTime: time.Date(2017, 7, 22, 19, 0, 0, 0, time.Local).UTC(), //  SAT | JULY 22, 2017 | 7:00 PM PT
	},
	SoccerGame{
		HomeTeam:  "Courage",
		AwayTeam:  "Reign",
		StartTime: time.Date(2017, 8, 5, 16, 30, 0, 0, time.Local).UTC(), //  AUGUST 5, 2017 | 4:30 PM PT
	},
	SoccerGame{
		HomeTeam:  "Reign",
		AwayTeam:  "Courage",
		StartTime: time.Date(2017, 8, 13, 18, 0, 0, 0, time.Local).UTC(), //  SUN | AUGUST 13, 2017 | 6:00 PM PT
	},
	SoccerGame{
		HomeTeam:  "Red Stars",
		AwayTeam:  "Reign",
		StartTime: time.Date(2017, 8, 16, 17, 0, 0, 0, time.Local).UTC(), //  WED | AUGUST 16, 2017 | 5:00 PM PT
	},
	SoccerGame{
		HomeTeam:  "Sky Blue",
		AwayTeam:  "Reign",
		StartTime: time.Date(2017, 8, 19, 16, 0, 0, 0, time.Local).UTC(), //  SAT | AUGUST 19, 2017 | 4:00 PM PT
	},
	SoccerGame{
		HomeTeam:  "Reign",
		AwayTeam:  "Thorns",
		StartTime: time.Date(2017, 8, 26, 13, 0, 0, 0, time.Local).UTC(), //  SAT | AUGUST 26, 2017 | 1:00 PM PT
	},
	SoccerGame{
		HomeTeam:  "Dash",
		AwayTeam:  "Reign",
		StartTime: time.Date(2017, 9, 3, 17, 0, 0, 0, time.Local), //  SUN | SEPTEMBER 3, 2017 | 5:00 PM PT
	},
	SoccerGame{
		HomeTeam:  "Pride",
		AwayTeam:  "Reign",
		StartTime: time.Date(2017, 9, 9, 16, 30, 0, 0, time.Local), //  SAT | SEPTEMBER 9, 2017 | 4:30 PM PT
	},
	SoccerGame{
		HomeTeam:  "Reign",
		AwayTeam:  "Kansas",
		StartTime: time.Date(2017, 9, 24, 18, 0, 0, 0, time.Local), //  SUN | SEPTEMBER 24, 2017 | 6:00 PM PT
	},
	SoccerGame{
		HomeTeam:  "Spirit",
		AwayTeam:  "Reign",
		StartTime: time.Date(2017, 9, 30, 13, 0, 0, 0, time.Local), //  SAT | SEPTEMBER 30, 2017 | 1:00 PM PT
	},
}

func getReignSchedule() []SoccerGame {

	return reignGames
}
