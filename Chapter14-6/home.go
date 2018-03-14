package main

import (
	"fmt"
	"net/http"
	//"github.com/gorilla/mux"
)

func HomePageHandle(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func UsersHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Users Page")
}

func main() {
	http.HandleFunc("/", HomePageHandle)
	http.HandleFunc("/user", UsersHandle)
	http.ListenAndServe(":3000", nil)
}
