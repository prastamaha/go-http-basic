package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Welcome to my Site :D</h1>")
	})

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Nice to meet you!</h1>")
	})

	http.ListenAndServe(":8080", nil)
}
