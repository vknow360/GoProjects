package main

import (
	"fmt"
	"io"
	"os"

	"net/http"

	"github.com/tidwall/gjson"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <city> [-f]")
		return
	}
	city := os.Args[1]
	useFahrenheit := false

	if len(os.Args) > 2 && os.Args[2] == "-f" {
		useFahrenheit = true
	}
	getWeather(city, useFahrenheit)
}

func getWeather(city string, fahrenheit bool) {
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		fmt.Println("❌ API key not found! Set WEATHER_API_KEY environment variable.")
		return
	}
	res, err := http.Get(`http://api.weatherapi.com/v1/current.json?key=` + apiKey + `&q=` + city)
	if err != nil {
		fmt.Println("❌ Error making request:", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("❌ Error reading response:", err)
		return
	}

	if gjson.GetBytes(body, "error").Exists() {
		fmt.Println("⚠️ Invalid city name! Please enter a valid location.")
		return
	}

	location := gjson.GetBytes(body, "location.name").String()
	country := gjson.GetBytes(body, "location.country").String()
	temperature := gjson.GetBytes(body, "current.temp_c").Float()
	tempUnit := " °C"
	if fahrenheit {
		temperature = gjson.GetBytes(body, "current.temp_f").Float()
		tempUnit = " °F"
	}
	feelsLike := gjson.GetBytes(body, "current.feelslike_c").Float()
	if fahrenheit {
		feelsLike = gjson.GetBytes(body, "current.feelslike_f").Float()
	}
	condition := gjson.GetBytes(body, "current.condition.text").String()
	windSpeed := gjson.GetBytes(body, "current.wind_kph").Float()
	humidity := gjson.GetBytes(body, "current.humidity").Int()

	fmt.Printf("🌍 Weather in %s, %s:\n", location, country)
	fmt.Printf("🌡️  Temperature: %.1f%s\n", temperature, tempUnit)
	fmt.Printf("🥶 Feels Like: %.1f%s\n", feelsLike, tempUnit)
	fmt.Printf("🌤️  Condition: %s\n", condition)
	fmt.Printf("💨 Wind Speed: %.1f km/h\n", windSpeed)
	fmt.Printf("💧 Humidity: %d%%\n", humidity)
}
