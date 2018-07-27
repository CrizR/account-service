package server

import (
	"encoding/json"
	"net/http"
)

func (s *server) getRoutes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Is there a better way to return this?
	var endpoints struct {
		PrintRoutes    string `json:"/api/"`
		CreateUser     string `json:"/api/create/{userId}"`
		GetUserByID    string `json:"/api/search/{userId}"`
		GetAllUsers    string `json:"/api/search/users"`
		GetUserByEmail string `json:"/api/search/{email}"`
		RemoveUserByID string `json:"/api/remove/{userId}"`
		UpdateUser     string `json:"/api/update/{userId}"`
	}
	json.NewEncoder(w).Encode(endpoints)
}

func (s *server) createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func (s *server) getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func (s *server) getUserByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func (s *server) getUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func (s *server) removeUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func (s *server) updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}
