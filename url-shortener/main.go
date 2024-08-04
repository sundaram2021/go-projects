package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var (
	urlStore = make(map[string]string)
	mutex    = &sync.Mutex{}
)

func main() {
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/", redirectHandler)

	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	longURL := r.URL.Query().Get("url")
	if longURL == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}

	shortURL := generateShortURL(longURL)

	mutex.Lock()
	urlStore[shortURL] = longURL
	mutex.Unlock()

	saveURLMapping()

	w.Write([]byte(fmt.Sprintf("Short URL: http://localhost:8080/%s", shortURL)))
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:]

	mutex.Lock()
	longURL, exists := urlStore[shortURL]
	mutex.Unlock()

	if !exists {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}

func generateShortURL(longURL string) string {
	// Implement a method to generate a short URL, like using a hash or a counter
	return fmt.Sprintf("%x", len(longURL)) // Simplified example
}

func saveURLMapping() {
	data, err := json.Marshal(urlStore)
	if err != nil {
		log.Println("Error marshaling data:", err)
		return
	}

	err = ioutil.WriteFile("urls.json", data, 0644)
	if err != nil {
		log.Println("Error writing to file:", err)
	}
}

func loadURLMapping() {
	data, err := ioutil.ReadFile("urls.json")
	if err != nil {
		log.Println("Error reading file:", err)
		return
	}

	err = json.Unmarshal(data, &urlStore)
	if err != nil {
		log.Println("Error unmarshaling data:", err)
	}
}
