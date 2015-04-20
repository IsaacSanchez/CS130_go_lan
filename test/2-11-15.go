// This program creates a web app that runs differnt functions for corresponding URL paths

package main

import (
	"fmt"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the default handler")
}

func oneHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You are using the first Handler")
}

func twoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You are using the second Handler")
}

func threeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You are using the third handler")
}

func main() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/test1/", oneHandler)
	http.HandleFunc("/test2/", twoHandler)
	http.HandleFunc("/test3/", threeHandler)
	http.ListenAndServe(":8080", nil)
}
