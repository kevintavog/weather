package main

import (
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/ian-kent/go-log/log"
)

var HomeDirectory string
var DataDirectory string
var LogDirectory string
var ConfigDirectory string
var ExecutingDirectory string

func InitDirectories() {
	if HomeDirectory == "" {
		exeDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal("Unable to get executing directory: %s", err)
		}

		ExecutingDirectory = exeDir

		if runtime.GOOS == "darwin" {
			HomeDirectory = os.Getenv("HOME")
			DataDirectory = path.Join(HomeDirectory, "Library", "Caches", "rangic.weather")
			LogDirectory = path.Join(HomeDirectory, "Library", "Logs", "weather")
			ConfigDirectory = path.Join(HomeDirectory, "Library", "Preferences")
		} else if runtime.GOOS == "linux" {
			HomeDirectory = os.Getenv("HOME")
			DataDirectory = path.Join(HomeDirectory, ".weather", "cache")
			LogDirectory = path.Join(HomeDirectory, ".weather", "logs")
			ConfigDirectory = path.Join(HomeDirectory, ".weather")
		} else {
			log.Fatal("Come up with directories for: %v", runtime.GOOS)
		}

		err = os.MkdirAll(DataDirectory, os.ModePerm)
		if err != nil {
			log.Fatal("Unable to create data directory (%s): %s", DataDirectory, err.Error())
		}

	}
}
