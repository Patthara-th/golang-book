package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"strconv"
)

type Weatherresult struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`

	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`

	Base string `json:"base"`

	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int `json:"pressure"`
		Humidity int `json:"humidity"`
		TempMin  int `json:"temp_min"`
		TempMax  int `json:"temp_max"`
	} `json:"main"`

	Visibility int `json:"visibility"`

	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`

	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`

	Dt  int `json:"dt"`

	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}

func HomePageHandle(w http.ResponseWriter, r *http.Request) {
	// start := time.Now()
	// fmt.Printf("Start at %v", start)

	vars := mux.Vars(r)
	name := vars["name"]
	var result string

	if name == "all" {
		result = getallcity()
	} else {
		result = getcity(name)
	}

	fmt.Fprintf(w, result)

	// fmt.Printf("Completed in %v", time.Since(start))
}

func getallcity() string {
	city := [...]string{"hobart","newyork","kupang","nairobi","bangkok"}

	
	var result string
	//result1 := make(chan string)

	for _,v := range city {
		result += getcity(v)
	}		

	return result
}

func getcity(city string) string {

	URL := "http://localhost:8882/api/v1/weather/"
	res, _ := http.Get(URL + city)
	
	weather := new(Weatherresult)
	json.NewDecoder(res.Body).Decode(weather)	
	
	return weather.Name  + "\n" + strconv.FormatFloat(weather.Main.Temp, 'f', 0, 64) +  "c " + weather.Weather[0].Description + "\n\n"
}


func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("Start at %v", start)
		next.ServeHTTP(w, r)
		fmt.Printf(" -> Completed in %v\n", time.Since(start))
	})
}

func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/weather/{name}", HomePageHandle).Methods("GET")	
	r.Use(loggingMiddleware)
	return r
}

func main() {
	var port string

	flag.StringVar(&port, "port", ":3000", "default port: 3000")
	flag.Parse()

	http.ListenAndServe(port, NewRouter())
}
