package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func HomePageHandle(w http.ResponseWriter, r *http.Request) {
	// start := time.Now()
	// fmt.Printf("Start at %v", start)

	vars := mux.Vars(r)
	name := vars["name"]
	if vars["name"] == "" {
		name = "World"
	}

	fmt.Fprintf(w, "Hello, %s!", name)

	// fmt.Printf("Completed in %v", time.Since(start))
}
func UsersHandle(w http.ResponseWriter, r *http.Request) {
	// start := time.Now()
	// fmt.Printf("Start at %v", start)

	fmt.Fprintf(w, "Users Page")

	// fmt.Printf("Completed in %v", time.Since(start))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("Start at %v", start)
		next.ServeHTTP(w, r)
		fmt.Printf("Completed in %v", time.Since(start))
	})
}

func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/{name}", HomePageHandle).Methods("GET")
	r.HandleFunc("/", UsersHandle).Methods("GET")
	r.Use(loggingMiddleware)
	return r
}

func main() {
	var port string

	flag.StringVar(&port, "port", ":3000", "default port: 3000")
	flag.Parse()

	http.ListenAndServe(port, NewRouter())
}

// go build -o main.exe
// main.exe -port=:8080
// .\main.exe -port=:8080

// func HomePageHandle(w http.ResponseWriter, r *http.Request) {
// 	name := r.URL.Query().Get("name")
// 	if name == "" {
// 		name = "World"
// 	}
// 	fmt.Fprintf(w, "Hello, %s!", name)
// }

// func UsersHandle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Users Page")
// }

// func main() {
// 	http.HandleFunc("/", HomePageHandle)
// 	http.HandleFunc("/user", UsersHandle)
// 	http.ListenAndServe(":3000", nil)
// }
