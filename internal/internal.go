package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/karlo235/DHBW-DevOps/pkg"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func AuthLoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username := r.FormValue("username")
	password := r.FormValue("password")

	// For simplicity, we'll use a hardcoded username and password
	// This should be replaced with a proper authentication mechanism
	if username == "user" && password == "pass" {
		token, err := createToken(username)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Error generating the token"}`))
			return
		}

		w.Write([]byte(fmt.Sprintf(`{"token": "%s"}`, token)))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "Invalid credentials"}`))
	}
}

func AuthLogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// In this simple example, we'll just return a success message
	w.Write([]byte(`{"message": "Logout successful"}`))
}

func FindProductByID(products []Product, id int) *Product {
	for _, product := range products {
		if product.ID == id {
			return &product
		}
	}
	return nil
}

func ProductListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Internal Server Error"}`))
		return
	}

	w.Write(response)
}

func ProductDetailHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	productID, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Product ID is missing"}`))
		return
	}

	id, err := strconv.Atoi(productID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Product ID has wrong format"}`))
		return
	}

	product := findProductByID(products, id)
	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Product not found"}`))
		return
	}

	response, err := json.Marshal(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Internal Server Error"}`))
		return
	}

	w.Write(response)
}

func CheckoutPlaceOrderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := r.Header.Get("Authorization")

	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "Missing authorization header"}`))
		return
	}

	if verifyToken(token) {
		// In this simple example, we'll just return a success message
		w.Write([]byte(`{"message": "Order placed successfully"}`))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "Invalid token"}`))
	}
}