package main

import (
	// "encoding/json"
	// "log"
	"net/http"
	// "reflect"
)



func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	http.ListenAndServe(":8080", nil)
}