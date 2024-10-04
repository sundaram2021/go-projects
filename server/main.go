package main

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

type Authentication struct {
    Password     string `json:"password"`
    Salt         string `json:"salt"`
    SessionToken string `json:"sessionToken"`
}

type UserLogin struct {
    ID             int32           `json:"id"`
    Name           string          `json:"name"`
    Authentication Authentication `json:"authentication"`
}

type UserRegister struct {
    ID       int32  `json:"id"`
    Name     string `json:"name"`
    Password string `json:"password"`
}

var LoggedUsers []UserLogin
var RegisteredUsers []UserRegister

func respondWithError(w http.ResponseWriter, statusCode int, message string) {
    w.WriteHeader(statusCode)
    fmt.Fprintf(w, "Message: %v", message)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to gorilla mux, URL: %v", r.RequestURI)
}

func numHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    fmt.Fprintf(w, "Requested Number is: %v", vars["num"])
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query().Get("query")
    fmt.Fprintf(w, "Query parameter is: %v", query)
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

func registerHandler(w http.ResponseWriter, r *http.Request) {
    var data UserRegister
    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        respondWithError(w, http.StatusBadRequest, "JSON encoding error while registering")
        return
    }
    RegisteredUsers = append(RegisteredUsers, data)
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "User is registered: %+v", data)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    var data UserLogin
    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        respondWithError(w, http.StatusBadRequest, "Error in JSON encoding while logging in")
        return
    }
    LoggedUsers = append(LoggedUsers, data)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "User is logged in: %v", data.Name)
}

func main() {
    mux := mux.NewRouter()

    mux.HandleFunc("/", rootHandler).Methods("GET")
    mux.HandleFunc("/{num}", numHandler).Methods("GET")
    mux.HandleFunc("/query", queryHandler).Methods("GET")  // Specific path for query
    mux.HandleFunc("/login", loginHandler).Methods("POST")
    mux.HandleFunc("/register", registerHandler).Methods("POST")

    fmt.Println("Mux Server is running on port 8080...")
    http.ListenAndServe(":8080", mux)
}
