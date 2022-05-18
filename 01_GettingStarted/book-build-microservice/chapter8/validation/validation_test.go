package validation

import (
	"testing"

	"gopkg.in/go-playground/validator.v9"
)
/*******************************************************************************
Test case for structure email
*******************************************************************************/
// Case 1: Test return error if a email is not exist in a request
func TestErrorWhenRequestEmailNotPresent(t *testing.T) {
	validate := validator.New()
	request := Request{
		URL: "http://nicholasjackson.io",
	}

	// err == nil: validate return error
	if err := validate.Struct(&request); err == nil {
		t.Error("Should have raised an error")
	}
}

// Case 2: Test return error if a email is invalid
func TestErrorWhenRequestEmailIsInvalid(t *testing.T) {
	validate := validator.New()
	request := Request{
		Email: "something.com",
		URL:   "http://nicholasjackson.io",
	}

	// err == nil: validate return error
	if err := validate.Struct(&request); err == nil {
		t.Error("Should have raised an error")
	}
}

// Case 3: Test return success if a name is not present
func TestNoErrorWhenRequestNameNotPresent(t *testing.T) {
	validate := validator.New()
	request := Request{
		Email: "myname@address.com",
		URL:   "http://nicholasjackson.io",
	}

	// err != nil: validate return ok
	if err := validate.Struct(&request); err == nil {
		t.Error(err)
	}
}

// Case 4: Test return error if a name is present but a length is in range
func TestNoErrorWhenRequestNamePresent(t *testing.T) {
	validate := validator.New()
	request := Request{
		Name:  "Hua Van Thong",
		Email: "myname@address.com",
		URL:   "http://nicholasjackson.io",
	}

	// err != nil: validate return ok
	if err := validate.Struct(&request); err != nil {
		t.Error(err)
	}
}

// Case 5: Test return error if a name is present but a length is greater 30
func TestNoErrorWhenRequestNamePresentGreater(t *testing.T) {
	validate := validator.New()
	request := Request{
		Name:  "Hua Van Thong Hackerrrrrrrrrrrrrrrrrrrrrr",
		Email: "myname@address.com",
		URL:   "http://nicholasjackson.io",
	}

	// err != nil: validate return ok
	if err := validate.Struct(&request); err == nil {
		t.Error(err)
	}
}

/*******************************************************************************
Test case for structure weather data
*******************************************************************************/
// Case 1: Normal teset
func TestNoErrorWithNormalWeatherData(t *testing.T) {
	validate := validator.New()

	weatherData := WeatherData {
		Index : 1,
		Coord: Coordinate{
			Latitude: 8.85023121875,
			Longitude: 103.904589453125,
		},
		Country: "Viet Nam",
	}


	// err == nil: validate return error
	if err := validate.Struct(&weatherData); err != nil {
		t.Error("Should have raised an error")
	}
}
