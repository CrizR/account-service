package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

var bindTo string


func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/", getRoutes).Methods("GET", "POST")
	router.HandleFunc("/api/create/{userId}", createUser).Methods("POST")
	router.HandleFunc("/api/search/{userId}", getUserById).Methods("POST")
	router.HandleFunc("/api/search/users", getAllUsers).Methods("GET", "POST")
	router.HandleFunc("/api/search/{email}", getUserByEmail).Methods("POST")
	router.HandleFunc("/api/remove/{userId}", removeUserById).Methods("POST")
	router.HandleFunc("/api/update/{userId}", updateUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}


func getRoutes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Is there a better way to return this?
	var endpoints struct {
		PrintRoutes string `json:"/api/"`
		CreateUser string `json:"/api/create/{userId}"`
		GetUserById string `json:"/api/search/{userId}"`
		GetAllUsers string `json:"/api/search/users"`
		GetUserByEmail string `json:"/api/search/{email}"`
		RemoveUserById string `json:"/api/remove/{userId}"`
		UpdateUser string `json:"/api/update/{userId}"`
	}
	json.NewEncoder(w).Encode(endpoints);
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")


}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")


}

func getUserByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func getUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")


}

func removeUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")


}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")


}