package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

var bindTo string

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/", printRoutes).Methods("GET", "POST")
	router.HandleFunc("/api/create/{userId}", createUser).Methods("POST")
	router.HandleFunc("/api/search/{userId}", getUserById).Methods("POST")
	router.HandleFunc("/api/search/users", getAllUsers).Methods("GET", "POST")
	router.HandleFunc("/api/", getUserByEmail).Methods("POST")
	router.HandleFunc("/api/", removeUserById).Methods("POST")
	router.HandleFunc("/api/update/{userId}", updateUser).Methods("POST")
	err := http.ListenAndServe(bindTo, nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}


func printRoutes(w http.ResponseWriter, r *http.Request) {

}

func createUser(w http.ResponseWriter, r *http.Request) {

}

func getAllUsers(w http.ResponseWriter, r *http.Request) {

}

func getUserByEmail(w http.ResponseWriter, r *http.Request) {

}

func getUserById(w http.ResponseWriter, r *http.Request) {

}

func removeUserById(w http.ResponseWriter, r *http.Request) {

}

func updateUser(w http.ResponseWriter, r *http.Request) {

}