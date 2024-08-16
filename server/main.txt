package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	Id   uint64
	Name string
}

var Users []User

func createUserFunc(id uint64, name string) User {
	usr := User{
		Id:   id,
		Name: name,
	}

	Users = append(Users, usr)

	return usr
}

func getUserFunc(id uint64) User {
	var usr User

	for _, ele := range Users {
		if ele.Id == id {
			usr = ele
			break
		}
	}

	return usr
}

func deleteUserFunc(id uint64) User {
	var usr User

	for i, ele := range Users {
		if ele.Id == id {
			usr = Users[i]
			Users = append(Users[:i], Users[i+1:]...)
			break
		}
	}

	return usr
}

func updateUserFunc(usr User) User {
	var usr2 User

	for i, ele := range Users {
		if ele.Id == usr.Id { // Corrected this line
			Users[i] = usr
			usr2 = Users[i]
			break
		}
	}

	return usr2
}

func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, "Message: %v", message)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusBadRequest, "Error, Bad Request")
		return
	}
	respondWithError(w, http.StatusOK, "Hello, welcome to the server")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var newUser User

		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		if newUser.Id == 0 || newUser.Name == "" {
			respondWithError(w, http.StatusNoContent, "ID or name is missing")
			return
		}

		createUserFunc(newUser.Id, newUser.Name)

		respondWithError(w, http.StatusCreated, "User created successfully")
		return
	}
	respondWithError(w, http.StatusBadRequest, "Bad Request")
}

func getAllUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := json.NewEncoder(w).Encode(Users)
		if err != nil {
			respondWithError(w, http.StatusBadGateway, "Encoding error")
			return
		}
		return
	}
	respondWithError(w, http.StatusBadRequest, "Bad Request")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		var usr User

		err := json.NewDecoder(r.Body).Decode(&usr)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		for _, ele := range Users {
			if ele.Id == usr.Id {
				deleteUserFunc(ele.Id)
				respondWithError(w, http.StatusOK, "User deleted successfully")
				return
			}
		}

		respondWithError(w, http.StatusNotFound, "User not found")
		return
	}
	respondWithError(w, http.StatusBadRequest, "Bad Request")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		var usr User

		err := json.NewDecoder(r.Body).Decode(&usr)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		updatedUser := updateUserFunc(usr)
		if updatedUser.Id == 0 {
			respondWithError(w, http.StatusNotFound, "User not found")
		} else {
			respondWithError(w, http.StatusOK, "User updated successfully")
		}
		return
	}
	respondWithError(w, http.StatusBadRequest, "Bad Request")
}

func getUser(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        // Extract the path segment to get the ID
		pathParts := strings.Split(r.URL.Path, "/")
		if len(pathParts) < 3 || pathParts[2] == "" {
			respondWithError(w, http.StatusBadRequest, "ID is missing in URL path")
			return
		}

		// Convert the ID from string to uint64
		idStr := pathParts[2]
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid ID format")
			return
		}

        // Fetch the user with the given ID
        user := getUserFunc(id)
        if user.Id == 0 {
            respondWithError(w, http.StatusNotFound, "User not found")
            return
        }
    }

    respondWithError(w, http.StatusBadRequest, "Bad Request")
}

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/create", createUser)
	r.HandleFunc("/getAllUser", getAllUser)
    r.HandleFunc("/getUser/{id}", getUser)
	r.HandleFunc("/deleteUser", deleteUser)
	r.HandleFunc("/updateUser", updateUser)

	http.ListenAndServe(":80", r)
}
