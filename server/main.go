package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

type Authentication struct {
	password string `json:password`
	salt string `json:salt`
	sessionToken string `json:sessionToken`
}
type UserLogin struct {
	id int32 `json:id`
	name string `json:name`
	authentication Authentication
}


type UserRegister struct {
	id int32 `json:id`
	name string `json:name`
	password string `json:password`
	
}

var LoggedUsers []UserLogin
var RegisteredUsers []UserRegister


func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, "Message: %v", message)
}

func rootHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to gorilla mux , URL : %v", r.RequestURI)
}

func numHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	fmt.Fprintf(w, "Requested Number is : %v", vars["num"])
}

func queryHandler(w http.ResponseWriter, r *http.Request){
	queryParams := r.URL.Query()

	query := queryParams.Get("query")

	fmt.Fprintf(w, "Query parameter is : %v", query)
}

func RouteMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(r.URL.Query()) > 0 {
			queryHandler(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func registerHandler(w http.ResponseWriter, r *http.Request){
	var data UserRegister

	err:= json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		respondWithError(w, http.StatusBadGateway, "json encoding error while registering")
	}

	RegisteredUsers = append(RegisteredUsers, data)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User is registered %v", data)
}

func loginHandler(w http.ResponseWriter, r *http.Request){
	var data reflect.Type

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		respondWithError(w, http.StatusBadGateway, "error in json encoding while logging")
	}
	LoggedUsers = append(LoggedUsers, data)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User is logged in %v", data.name)
}

func main() {
	
	mux := mux.NewRouter()

	mux.Handle("/" , RouteMiddleware(http.HandlerFunc(rootHandler))).Methods("GET").Schemes("http")
	mux.HandleFunc("/{num}", numHandler).Methods("GET").Schemes("http")
	mux.Handle("/", RouteMiddleware(http.HandlerFunc(queryHandler))).Methods("GET").Schemes("http")
	mux.HandleFunc("/login", loginHandler).Methods("POST").Schemes("http")
	mux.HandleFunc("/register", registerHandler).Methods("POST").Schemes("http")
	fmt.Println("Mux Server is running on port 8080...") 
	http.ListenAndServe(":8080", mux)

}