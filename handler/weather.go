package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rafaeldias/async"
)

type DarkSkyTimeMachineResponse struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	TimeZone  string  `json:"timezone"`
	Hourly    struct {
		Data []struct {
			Time                     int64   `json:"time"`
			Icon                     string  `json:"icon"`
			PrecipitationProbability float32 `json:"precipProbability"`
			Temperature              float32 `json:"temperature"`
			ApparentTemperature      float32 `json:"apparentTemperature"`
			WindSpeed                float32 `json:"windSpeed"`
		} `json:"data"`
	} `json:"hourly"`
	Daily struct {
		Data []struct {
			Time                     int64   `json:"time"`
			Summary                  string  `json:"summary"`
			Icon                     string  `json:"icon"`
			HighTemperature          float32 `json:"temperatureHigh"`
			ApparentHighTemperature  float32 `json:"apparentTemperatureHigh"`
			LowTemperature           float32 `json:"temperatureLow"`
			ApparentlowTemperature   float32 `json:"apparentTemperatureLow"`
			PrecipitationProbability float32 `json:"precipProbability"`
			WindSpeed                float32 `json:"windSpeed"`
			Humidity                 float32 `json:"humidity"`
		} `json:"data"`
	} `json:"daily"`
}

type DarkSkyForecastResponse struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	TimeZone  string  `json:"timezone"`
	Currently struct {
		Time                     int64   `json:"time"`
		Summary                  string  `json:"summary"`
		Icon                     string  `json:"icon"`
		PrecipitationProbability float32 `json:"precipProbability"`
		Temperature              float32 `json:"temperature"`
		ApparentTemperature      float32 `json:"apparentTemperature"`
		Humidity                 float32 `json:"humidity"`
		WindSpeed                float32 `json:"windSpeed"`
	} `json:"currently"`
	Hourly struct {
		Data []struct {
			Time                     int64   `json:"time"`
			Icon                     string  `json:"icon"`
			PrecipitationProbability float32 `json:"precipProbability"`
			Temperature              float32 `json:"temperature"`
			ApparentTemperature      float32 `json:"apparentTemperature"`
			WindSpeed                float32 `json:"windSpeed"`
		} `json:"data"`
	} `json:"hourly"`
	Daily struct {
		Summary string `json:"summary"`
		Data    []struct {
			Time                     int64   `json:"time"`
			Summary                  string  `json:"summary"`
			Icon                     string  `json:"icon"`
			HighTemperature          float32 `json:"temperatureHigh"`
			ApparentHighTemperature  float32 `json:"apparentTemperatureHigh"`
			LowTemperature           float32 `json:"temperatureLow"`
			ApparentlowTemperature   float32 `json:"apparentTemperatureLow"`
			PrecipitationProbability float32 `json:"precipProbability"`
			WindSpeed                float32 `json:"windSpeed"`
			Humidity                 float32 `json:"humidity"`
		} `json:"data"`
	} `json:"daily"`
}

type DarkSkyResponse struct {
	Historic []DarkSkyTimeMachineResponse `json:"historic"`
	Forecast DarkSkyForecastResponse      `json:"forecast"`
}

type Weather struct {
}

func (w Weather) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	go fmt.Print("Request Received.")
	vars := mux.Vars(request)
	dsr, err := w.getAllWeather(vars["latitude"], vars["longitude"])
	if err != nil {
		writer.Write([]byte(err.Error()))
	}
	body, err := json.Marshal(dsr)
	if err != nil {
		writer.Write([]byte(err.Error()))
	}
	writer.Write(body)
}

func (w Weather) getTimeMachineWeather(lat string, long string, time int64) (DarkSkyTimeMachineResponse, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.darksky.net/forecast//%s,%s,%d", lat, long, time))
	if err != nil {
		return DarkSkyTimeMachineResponse{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return DarkSkyTimeMachineResponse{}, err
	}
	var dstmr DarkSkyTimeMachineResponse
	err = json.Unmarshal(body, &dstmr)
	if err != nil {
		return DarkSkyTimeMachineResponse{}, err
	}

	return dstmr, nil
}

func (w Weather) getForecastWeather(lat string, long string) (DarkSkyForecastResponse, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.darksky.net/forecast//%s,%s", lat, long))
	if err != nil {
		return DarkSkyForecastResponse{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return DarkSkyForecastResponse{}, err
	}
	var dsfr DarkSkyForecastResponse
	err = json.Unmarshal(body, &dsfr)
	if err != nil {
		return DarkSkyForecastResponse{}, nil
	}

	return dsfr, nil
}

func (w Weather) getAllWeather(lat string, long string) (DarkSkyResponse, error) {
	currentTime := time.Now().Unix()

	mapTasks := async.MapTasks{
		"day1": func() (DarkSkyTimeMachineResponse, error) {
			return w.getTimeMachineWeather(lat, long, currentTime-86400)
		},
		"day2": func() (DarkSkyTimeMachineResponse, error) {
			return w.getTimeMachineWeather(lat, long, currentTime-86400*2)
		},
		"day3": func() (DarkSkyTimeMachineResponse, error) {
			return w.getTimeMachineWeather(lat, long, currentTime-86400*3)
		},
		"day4": func() (DarkSkyTimeMachineResponse, error) {
			return w.getTimeMachineWeather(lat, long, currentTime-86400*4)
		},
		"day5": func() (DarkSkyTimeMachineResponse, error) {
			return w.getTimeMachineWeather(lat, long, currentTime-86400*5)
		},
		"day6": func() (DarkSkyTimeMachineResponse, error) {
			return w.getTimeMachineWeather(lat, long, currentTime-86400*6)
		},
		"day7": func() (DarkSkyTimeMachineResponse, error) {
			return w.getTimeMachineWeather(lat, long, currentTime-86400*7)
		},
		"forecast": func() (DarkSkyForecastResponse, error) {
			return w.getForecastWeather(lat, long)
		},
	}
	results, err := async.Parallel(mapTasks)
	if err != nil {
		return DarkSkyResponse{}, err
	}

	var historic = make([]DarkSkyTimeMachineResponse, 7)
	historic[0] = results.Key("day7")[0].(DarkSkyTimeMachineResponse)
	historic[1] = results.Key("day6")[0].(DarkSkyTimeMachineResponse)
	historic[2] = results.Key("day5")[0].(DarkSkyTimeMachineResponse)
	historic[3] = results.Key("day4")[0].(DarkSkyTimeMachineResponse)
	historic[4] = results.Key("day3")[0].(DarkSkyTimeMachineResponse)
	historic[5] = results.Key("day2")[0].(DarkSkyTimeMachineResponse)
	historic[6] = results.Key("day1")[0].(DarkSkyTimeMachineResponse)

	//fmt.Printf("%+v\n", historic)

	dsr := DarkSkyResponse{
		Historic: historic,
		Forecast: results.Key("forecast")[0].(DarkSkyForecastResponse),
	}
	dsr.Historic = historic
	return dsr, nil
}
