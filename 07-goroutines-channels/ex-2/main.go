package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Get API Key from the https://openweathermap.org/ and locate in .env file
// http://api.openweathermap.org/data/2.5/weather?q={CITY}&appid={API_KEY}

type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
}

type WeatherResult struct {
	Data WeatherResponse
	Err  error
}

func main() {
	godotenv.Load()
	apiKey := os.Getenv("OPENWEATHER_API_KEY")

	if apiKey == "" {
		fmt.Printf("Can not get API KEY")
		return
	}

	cities := []string{"Toronto", "London", "Paris", "Tokyo", "Istanbul", "Moscow", "Oslo", "Ankara"}

	ch := make(chan WeatherResult)

	startTime := time.Now()

	// Launch all goroutines
	for _, city := range cities {
		go fetchWeather(city, apiKey, ch)
	}

	// Collect results (one receive per goroutine)
	for range cities {
		result := <-ch
		if result.Err != nil {
			fmt.Println("Error:", result.Err)
			continue
		}
		fmt.Printf("City: %v, Temperature: %v\n", result.Data.Name, result.Data.Main.Temp)
	}

	fmt.Printf("Time taken to fetch all cities: %v", time.Since(startTime))
}

func fetchWeather(city string, apiKey string, ch chan<- WeatherResult) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		ch <- WeatherResult{Err: err} // send error through channel
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- WeatherResult{Err: err}
		return
	}

	var data WeatherResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		ch <- WeatherResult{Err: err}
		return
	}

	ch <- WeatherResult{Data: data} // send success through channel
}
