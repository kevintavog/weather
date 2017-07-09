package main

import (
	"os"
	"path"

	"github.com/ian-kent/go-log/appenders"
	"github.com/ian-kent/go-log/layout"
	"github.com/ian-kent/go-log/log"
)

func ConfigureLogging(logDirectory, appName string) {

	err := os.MkdirAll(logDirectory, os.ModePerm)
	if err != nil {
		log.Fatal("Unable to create logging directory (%s): %s", logDirectory, err.Error())
	}

	logger := log.Logger("")

	lyt := layout.Pattern("%d %p: %m")
	layout.DefaultTimeLayout = "15:04:05.000000"

	rolling := appenders.RollingFile(path.Join(logDirectory, appName+".log"), true)
	rolling.MaxBackupIndex = 10
	rolling.MaxFileSize = 5 * 1024 * 1024
	rolling.SetLayout(lyt)

	console := appenders.Console()
	console.SetLayout(lyt)

	logger.SetAppender(appenders.Multiple(lyt, rolling, console))
}
