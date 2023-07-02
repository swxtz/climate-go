package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")

	fmt.Println(apiKey)
	res, err := http.Get(apiUrl(apiKey, "DASADS"))

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {

		panic("Weather Api não esta disponivel")
	}

}

func apiUrl(apiKey string, city string) string {

	url := "http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no"

	if city != "" {
		formatedUrl := fmt.Sprintf(url, apiKey, city)
		return formatedUrl
	}

	formatedUrl := fmt.Sprintf(url, apiKey, "sao_paulo")
	return formatedUrl
}
