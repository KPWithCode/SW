package handlers;

import (
	"fmt"
	"net/http"
	"os"
	"encoding/json"
)

type WeatherResponse struct {
	Temperature float64 `json:"temp"`
	Description string  `json:"weather"`
}

func getWeather(location string) (WeatherResponse, error) {
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
    if apiKey == "" {
        return WeatherResponse{}, fmt.Errorf("OPENWEATHER_API_KEY environment variable not set")
    }

    apiUrl := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", location, apiKey)

    // Make a GET request to the OpenWeatherAPI endpoint
    resp, err := http.Get(apiUrl)
    if err != nil {
        return WeatherResponse{}, fmt.Errorf("failed to fetch weather data: %v", err)
    }
    defer resp.Body.Close()

    // Decode the JSON response
    var weatherResp WeatherResponse
    if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
        return WeatherResponse{}, fmt.Errorf("failed to decode weather response: %v", err)
    }

    return weatherResp, nil
}

func WeatherHandler(w http.ResponseWriter, r http.Request) {
	location := r.URL.Query().Get("location")
	if location == "" {
		http.Error(w, "location parmater is required", http.StatusBadRequest)
	}
	weatherResp, err := getWeather(location)
    if err != nil {
        http.Error(w, fmt.Sprintf("failed to fetch weather data: %v", err), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(weatherResp); err != nil {
        http.Error(w, fmt.Sprintf("failed to encode weather response: %v", err), http.StatusInternalServerError)
        return
    }
}