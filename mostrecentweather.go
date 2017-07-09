package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/go-playground/lars"
)

func mostRecentWeather(c lars.Context) {
	wc := c.(*WeatherContext)

	wc.FieldLogger.Time("mostRecentWeather", func() {

		response := make(map[string]interface{})
		fileInfo, count, err := getLatestWeather()
		propogateError(err, fmt.Sprintf("Unable to read from %s", weatherDataDirectory))

		response["totalDataFileCount"] = count
		if fileInfo != nil {
			response["timestamp"] = fileInfo.ModTime()

			fileData, err := ioutil.ReadFile(filepath.Join(weatherDataDirectory, fileInfo.Name()))
			propogateError(err, fileInfo.Name())

			darkSkyData := make(map[string]interface{})
			err = json.Unmarshal(fileData, &darkSkyData)
			propogateError(err, fmt.Sprintf("Failed parsing JSON from %s", fileInfo.Name()))
			response["data"] = darkSkyData
		}

		wc.WriteResponse(response)
	})
}
