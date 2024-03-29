package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/karlo235/DHBW-DevOps/auth-service/internal"
)

func main() {
	router := mux.NewRouter()

	// Auth Service
	router.HandleFunc("/auth/login", internal.AuthLoginHandler).Methods("POST")
	router.HandleFunc("/auth/logout", internal.AuthLogoutHandler).Methods("POST")

	port := 8080
	log.Printf("The server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}