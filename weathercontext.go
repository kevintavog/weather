package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/lars"
	"github.com/kevintavog/findaphoto/findaphotoserver/fieldlogger"
)

type WeatherContext struct {
	*lars.Ctx
	FieldLogger *fieldlogger.FieldLogger
}

func (wc *WeatherContext) Reset(w http.ResponseWriter, r *http.Request) {

	wc.Ctx.Reset(w, r)
	wc.FieldLogger = fieldlogger.New()
}

func (wc *WeatherContext) RequestComplete() {
	wc.FieldLogger.Add("url", wc.Ctx.Request().RequestURI)
	wc.FieldLogger.Add("method", wc.Ctx.Request().Method)
	wc.FieldLogger.Close(wc.Ctx.Response().Status())
}

func NewWeatherContext(l *lars.LARS) lars.Context {
	return &WeatherContext{
		Ctx:         lars.NewContext(l),
		FieldLogger: fieldlogger.New(),
	}
}

func (wc *WeatherContext) WriteResponse(m map[string]interface{}) {

	wc.Ctx.Response().Header().Set(lars.ContentType, lars.ApplicationJSON)
	wc.Ctx.Response().WriteString(wc.ToJson(m))
}

func (wc *WeatherContext) ToJson(m map[string]interface{}) string {
	return toJson(m, func(err error) {
		wc.Error(http.StatusInternalServerError, "JsonConversionFailed", "", err)
	})
}

func toJson(m map[string]interface{}, errorHandler func(error)) string {
	json, err := json.Marshal(m)
	if err != nil {
		errorHandler(err)
		return "{}"
	} else {
		return string(json)
	}
}

func (wc *WeatherContext) Error(httpCode int, errorCode, errorMessage string, err error) {
	wc.FieldLogger.Error(errorMessage, err)
	wc.FieldLogger.Add("errorCode", errorCode)

	wc.Ctx.Response().Header().Set(lars.ContentType, lars.ApplicationJSON)
	wc.Ctx.Response().WriteHeader(httpCode)

	data := map[string]interface{}{"errorCode": errorCode, "errorMessage": errorMessage}
	if err != nil {
		data["internalError"] = err.Error()
	}

	var ok = true
	json := toJson(data, func(err error) {
		wc.FieldLogger.Add("marhsalError", err.Error())
		wc.Ctx.Response().WriteString(fmt.Sprintf("{\"errorCode\":\"%s\"", errorCode))
	})
	if ok {
		wc.Ctx.Response().WriteString(string(json))
	}
}
