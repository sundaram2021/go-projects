package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Response structure
type Response struct {
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
}

type NameResponse struct {
	Name string `json:"name"`
}

// Helper function to parse float from query parameters
func parseFloatParam(param string) (float64, error) {
	value, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid parameter: %v", err)
	}
	return value, nil
}

// Add handler
func addHandler(w http.ResponseWriter, r *http.Request) {
	x, err1 := parseFloatParam(r.URL.Query().Get("x"))
	y, err2 := parseFloatParam(r.URL.Query().Get("y"))

	var response Response
	if err1 != nil || err2 != nil {
		response.Error = "Invalid input parameters"
	} else {
		response.Result = x + y
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Subtract handler
func subtractHandler(w http.ResponseWriter, r *http.Request) {
	x, err1 := parseFloatParam(r.URL.Query().Get("x"))
	y, err2 := parseFloatParam(r.URL.Query().Get("y"))

	var response Response
	if err1 != nil || err2 != nil {
		response.Error = "Invalid input parameters"
	} else {
		response.Result = x - y
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Multiply handler
func multiplyHandler(w http.ResponseWriter, r *http.Request) {
	x, err1 := parseFloatParam(r.URL.Query().Get("x"))
	y, err2 := parseFloatParam(r.URL.Query().Get("y"))

	var response Response
	if err1 != nil || err2 != nil {
		response.Error = "Invalid input parameters"
	} else {
		response.Result = x * y
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Divide handler
func divideHandler(w http.ResponseWriter, r *http.Request) {
	x, err1 := parseFloatParam(r.URL.Query().Get("x"))
	y, err2 := parseFloatParam(r.URL.Query().Get("y"))

	var response Response
	if err1 != nil || err2 != nil {
		response.Error = "Invalid input parameters"
	} else if y == 0 {
		response.Error = "Division by zero is not allowed"
	} else {
		response.Result = x / y
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Name handler for "/:name" route
func nameHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the name from the URL path
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 2 || pathParts[1] == "" {
		http.Error(w, "Name not provided", http.StatusBadRequest)
		return
	}
	name := pathParts[1]

	response := NameResponse{Name: name}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/subtract", subtractHandler)
	http.HandleFunc("/multiply", multiplyHandler)
	http.HandleFunc("/divide", divideHandler)

	// Handle name route (any path matching "/:name")
	// http.HandleFunc("/", nameHandler) // New route for capturing name from URL

	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
