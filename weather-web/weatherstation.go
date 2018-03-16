package main

import (
	"flag"
	"fmt"
	"github.com/Patthara-th/weather"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func HomePageHandle(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["name"]
	var result string

	if name == "all" {
		result = weather.Getallcity()
	} else {
		result = weather.Getcity(name)
	}

	fmt.Fprintf(w, result)

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
