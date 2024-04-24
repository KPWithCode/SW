package routes

import (
	"net/http"
	"sexyweather/handlers"
)

func WeatherRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		handlers.WeatherHandler(w, *r)
	})
}
