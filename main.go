package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/go-playground/lars"
	"github.com/ian-kent/go-log/log"
)

var debugMode = flag.Bool("d", false, "Debug mode")

func main() {
	var port = 3131

	flag.Parse()

	InitDirectories()
	ReadConfiguration()
	ConfigureLogging(LogDirectory, "weather")

	log.Info(
		"Weather backend, listening at http://localhost:%d, using location: %f, %f, polling stop: %s. Data store: %s",
		port, Current.Latitude, Current.Longitude, Current.OneBusAwayStop, DataDirectory)

	dataCollector()
	go dataPruner()
	startHttpServer(*debugMode, port)
}

func startHttpServer(easyExit bool, listenPort int) {

	contentDir := ExecutingDirectory + "/content"
	log.Info("Serving site content from %s", contentDir)
	httpContent := http.Dir(contentDir)
	_, e := httpContent.Open("index.html")
	if e != nil {
		log.Warn("Unable to get files from the '%s' folder: %s\n", contentDir, e.Error())
	}
	fs := http.FileServer(httpContent)

	l := lars.New()
	l.RegisterContext(NewWeatherContext)

	l.Get("/", fs)
	l.Get("/*", fs)
	ConfigureRouting(l)

	startServerFunc := func() {
		err := http.ListenAndServe(fmt.Sprintf(":%d", listenPort), l.Serve())
		if err != nil {
			log.Fatal("Failed starting the service: %s", err.Error())
		}
	}

	if easyExit {
		go startServerFunc()

		fmt.Println("Hit enter to exit")
		var input string
		fmt.Scanln(&input)
	} else {
		startServerFunc()
	}
}
