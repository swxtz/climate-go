package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")

	fmt.Println(apiKey)
	teste := apiUrl(apiKey, "")
	fmt.Println(teste)

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
