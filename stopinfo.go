package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/go-playground/lars"
)

func stopInfo(c lars.Context) {
	wc := c.(*WeatherContext)

	wc.FieldLogger.Time("stopInfo", func() {

		response := make(map[string]interface{})
		fileInfo, _, err := getLatestStopInfo()
		propogateError(err, fmt.Sprintf("Unable to read from %s", stopInfoDataDirectory))

		if fileInfo != nil {
			response["updatedTime"] = fileInfo.ModTime()
			response["timestamp"] = time.Now().Unix()

			fileData, err := ioutil.ReadFile(filepath.Join(stopInfoDataDirectory, fileInfo.Name()))
			propogateError(err, fileInfo.Name())

			oneBusAwayData := make(map[string]interface{})
			err = json.Unmarshal(fileData, &oneBusAwayData)
			propogateError(err, fmt.Sprintf("Failed parsing JSON from %s", fileInfo.Name()))
			response["data"] = oneBusAwayData
		}

		wc.WriteResponse(response)
	})
}
