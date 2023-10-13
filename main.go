package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Request struct {
	CardNumber int `json:"card_number"`
}

type ResponseStruct struct {
	Valid bool `json:"isValid"`
}

func main() {

	http.HandleFunc("/", validateHandler) // Create a route for validation

	log.Println("Server started at port 9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func validateHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method not suppported, Try GET", http.StatusNotFound)
		return
	}

	var requestPayload Request // Define your JSON structure here

	// Parse the JSON payload from the request body
	err := json.NewDecoder(r.Body).Decode(&requestPayload)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Implement validation logic here using the Luhn algorithm
	fmt.Println(requestPayload.CardNumber)
	isValid := Checker(requestPayload.CardNumber)

	// Create a JSON response
	response := ResponseStruct{
		Valid: isValid,
	}

	// Send the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// cardNumber := 4847352989263094
// if Checker(cardNumber) == true {
// 	fmt.Println("Valid Card")
// } else {
// 	fmt.Println("Not valid a card")
// }
