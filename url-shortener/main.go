package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

var (
	urlStore = make(map[string]string)
	mutex    = &sync.Mutex{}
)

func main() {
	// Load existing URL mappings from file
	loadURLMapping()

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
	saveURLMapping()
	mutex.Unlock()

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
	hash := sha1.New()
	hash.Write([]byte(longURL))
	return hex.EncodeToString(hash.Sum(nil))[:8] // Use the first 8 characters of the hash
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
	file, err := os.Open("urls.json")
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error reading file:", err)
		return
	}

	err = json.Unmarshal(data, &urlStore)
	if err != nil {
		log.Println("Error unmarshaling data:", err)
	}
}
