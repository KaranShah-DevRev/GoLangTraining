package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const GEOAPI = "http://api.openweathermap.org/geo/1.0/direct?q=%s&limit=1&appid=8be2d4ff1babea2b70deef5f1331d4f7"
const WEATHERAPI = "http://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=8be2d4ff1babea2b70deef5f1331d4f7&units=metric"

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type WeatherData struct {
	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Main struct {
		Temp    float64 `json:"temp"`
		TempMin float64 `json:"temp_min"`
		TempMax float64 `json:"temp_max"`
	} `json:"main"`
}

func getWeather(wg *sync.WaitGroup, city string) {
	defer wg.Done()
	resp, err := http.Get(fmt.Sprintf(GEOAPI, city))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	decoder := json.NewDecoder(resp.Body)

	var location []Location

	if err := decoder.Decode(&location); err != nil {
		fmt.Println(err)
		return
	}

	if len(location) == 0 {
		fmt.Println("No location found")
		return
	}

	lat := location[0].Lat
	lon := location[0].Lon

	fmt.Printf("Lat: %f, Lon: %f\n", lat, lon)

	resp, err = http.Get(fmt.Sprintf(WEATHERAPI, lat, lon))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	fmt.Println(resp.Status)
	decoder = json.NewDecoder(resp.Body)

	var weatherData WeatherData

	if err := decoder.Decode(&weatherData); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Temp: %f, TempMin: %f, TempMax: %f\n", weatherData.Main.Temp, weatherData.Main.TempMin, weatherData.Main.TempMax)
}

func main() {
	var wg sync.WaitGroup
	var cities = []string{"Chennai", "Jamshedpur", "Buenos Aires", "Tokyo"}
	var timePresent = time.Now()
	for _, city := range cities {
		wg.Add(1)
		go getWeather(&wg, city)
	}
	wg.Wait()
	var timEnd = time.Now()
	fmt.Println("Time taken: ", timEnd.Sub(timePresent))
}
