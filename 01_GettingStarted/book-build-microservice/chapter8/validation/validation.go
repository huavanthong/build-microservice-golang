package validation

import (
	"encoding/json"
	"net/http"

	validator "gopkg.in/go-playground/validator.v9"
)

// Request defines the input structure received by a http handler
type Request struct {
	Name  string `json:"name" validate:"required,gt=0,lt=30"`
	Email string `json:"email" validate:"email"`
	URL   string `json:"url" validate:"url"`
}

var validate = validator.New()

func Handler(rw http.ResponseWriter, r *http.Request) {
	request := Request{}

	err := json.NewEncoder(rw).Encode(&request)
	if err != nil {
		http.Error(rw, "Invalid request object", http.StatusBadRequest)
		return
	}

	err = validate.Struct(request)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

// Request defines the input structure of weather data

type Coordinate struct {
	Latitude  float32 `json:"latitude" validate:"latitude"`
	Longitude float32 `json:"longitude" validate:"longitude"`
}

type WeatherData struct {
	Index      int        `json:"id" validate:"required,gt=0"`
	coordinate Coordinate `json:"coordinate"`
	country    string     `json:"country"`
}

func HandlerWeatherData(rw http.ResponseWriter, r *http.Request) {
	wd := WeatherData{}

	err := json.NewEncoder(rw).Encode(&wd)
	if err != nil {
		http.Error(rw, "Invalid request object", http.StatusBadRequest)
		return
	}

	err = validate.Struct(wd)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusOK)
}
