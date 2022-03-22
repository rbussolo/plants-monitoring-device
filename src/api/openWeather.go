package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	OPEN_WEATHER_API_KEY   = "8492344b1bec905d88e7278368bb02fc"
	OPEN_WEATHER_FROM_CITY = "https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric"
)

type OpenData struct {
	Main struct {
		Temp     float32 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Dt   int64  `json:"dt"`
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func OpenWeatherGetDataFromCity(city string) (OpenData, error) {
	data := OpenData{}

	request_url := fmt.Sprintf(OPEN_WEATHER_FROM_CITY, city, OPEN_WEATHER_API_KEY)

	resp, err := http.Get(request_url)
	if err != nil {
		return data, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(body, &data)

	if err != nil {
		return data, err
	}

	return data, nil
}
