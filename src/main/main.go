package main

import (
	"fmt"
	"net/http"
)

func main() {
	// office code
	office_code := "MKX"

	x := 162
	y := 149

	// get forecast
	url := fmt.Sprintf("https://api.weather.gov/gridpoints/%s/%d/%d/forecast", office_code, x, y)
	resp, err := http.Get(url)

	if err != nil {
		// handle error
		fmt.Printf(err.Error())
	}
}
