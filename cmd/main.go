package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/karlo235/DHBW-DevOps/internal"
	"github.com/karlo235/DHBW-DevOps/pkg"
)

// Static data for three products
var products = []Product{
	{ID: 1, Name: "Product 1", Price: 19.99},
	{ID: 2, Name: "Product 2", Price: 29.99},
	{ID: 3, Name: "Product 3", Price: 39.99},
}

func main() {
	router := mux.NewRouter()

	// Auth Service
	router.HandleFunc("/auth/login", authLoginHandler).Methods("POST")
	router.HandleFunc("/auth/logout", authLogoutHandler).Methods("POST")

	// Product Service
	router.HandleFunc("/products", productListHandler).Methods("GET")
	router.HandleFunc("/products/{id}", productDetailHandler).Methods("GET")

	// Checkout Service
	router.HandleFunc("/checkout/placeorder", checkoutPlaceOrderHandler).Methods("POST")

	port := 8080
	log.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}