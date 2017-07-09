package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/go-playground/lars"
)

type errorApi struct {
	err     error
	message string
}

func (e *errorApi) Error() string {
	return e.message + " -- " + getDetailedErrorMessage(e.err)
}

func getDetailedErrorMessage(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

func ConfigureRouting(l *lars.LARS) {
	l.Use(handleErrors)

	api := l.Group("/api")
	api.Get("/stop", stopInfo)
	api.Get("/soccer", soccerScheduel)
	api.Get("/weather", mostRecentWeather)
}

func propogateError(err error, message string) {
	if err != nil {
		panic(&errorApi{message: message, err: err})
	}
}

func handleErrors(c lars.Context) {
	wc := c.(*WeatherContext)
	defer func() {
		if r := recover(); r != nil {

			logStack := true
			if e, ok := r.(runtime.Error); ok {
				wc.Error(http.StatusInternalServerError, "UnhandledError", "", e)
			} else if s, ok := r.(string); ok {
				wc.Error(http.StatusInternalServerError, "UnhandledError", s, nil)
			} else {
				wc.Error(http.StatusInternalServerError, "UnhandledError", fmt.Sprintf("%v", r), nil)
			}

			if logStack {
				buf := make([]byte, 1<<16)
				stackSize := runtime.Stack(buf, false)
				wc.FieldLogger.Add("stack", string(buf[0:stackSize]))
			}
		}
	}()

	wc.Ctx.Next()
}
