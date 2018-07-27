package server

import (
	"github.com/ecclesia-dev/account-service/controllers"
	"github.com/gorilla/mux"
)

type server struct {
	router   *mux.Router
	accounts controllers.AccountController
}

func (s *server) Start() {
	s.router = mux.NewRouter()
	s.setRoutes()
}

func (s *server) setRoutes() {
	s.router.HandleFunc("/api/", s.getRoutes).Methods("GET", "POST")
	s.router.HandleFunc("/api/create/{userId}", s.createUser).Methods("POST")
	s.router.HandleFunc("/api/search/{userId}", s.getUserByID).Methods("POST")
	s.router.HandleFunc("/api/search/users", s.getAllUsers).Methods("GET", "POST")
	s.router.HandleFunc("/api/search/{email}", s.getUserByEmail).Methods("POST")
	s.router.HandleFunc("/api/remove/{userId}", s.removeUserByID).Methods("POST")
	s.router.HandleFunc("/api/update/{userId}", s.updateUser).Methods("POST")
}
