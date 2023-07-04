package main

import (
	"fmt"

	"github.com/swxtz/sun/internal/config"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
		Hour    string `json:"localtime"`
	} `json:"location"`

	Current struct {
		TempC     float32 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
}

func main() {
	err := config.CreateConfigFile()

	if err != nil {
		fmt.Print(err)
		panic(err)
	}

	// 	res, err := http.Get(apiUrl("", ""))

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	defer res.Body.Close()

	// 	if res.StatusCode != 200 {

	// 		panic("Weather Api não esta disponivel")
	// 	}

	// 	body, err := io.ReadAll(res.Body)

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	var weather Weather
	// 	err = json.Unmarshal(body, &weather)

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	location, current := weather.Location, weather.Current

	// fmt.Printf("%s, %s, %.0f°C, %s \n", location.Name, location.Country, current.TempC, location.Hour)
}

// func apiUrl(apiKey string, city string) string {

// 	url := "http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no"

// 	if city != "" {
// 		formatedUrl := fmt.Sprintf(url, apiKey, city)
// 		return formatedUrl
// 	}

// 	formatedUrl := fmt.Sprintf(url, apiKey, "sao_paulo")
// 	return formatedUrl
// }
