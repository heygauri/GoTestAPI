package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

// Create
func addClient(newClient Client) error {
	for _, client := range clients {
		if client.ID == newClient.ID {
			return errors.New("client already exists")
		}
	}

	clients = append(clients, newClient)
	return nil
}

func createClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newClient Client
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // disallow unknown fields in the request body
	if err := decoder.Decode(&newClient); err != nil {
		// Handle errors in decoding request body
		w.WriteHeader(http.StatusBadRequest)
		message := "Invalid request body"
		json.NewEncoder(w).Encode(map[string]string{"error": message})
		return
	}

	// Check if client with given ID already exists
	for _, client := range clients {
		if client.ID == newClient.ID {
			w.WriteHeader(http.StatusBadRequest) // Set the status code to 400
			message := "Client with given ID already exists"
			json.NewEncoder(w).Encode(map[string]string{"error": message})
			return
		}
	}

	// Check if newClient has any unknown fields
	v := reflect.ValueOf(newClient)
	for i := 0; i < v.NumField(); i++ {
		fieldName := v.Type().Field(i).Name
		if !strings.HasPrefix(fieldName, "ID") &&
			!strings.HasPrefix(fieldName, "Name") &&
			!strings.HasPrefix(fieldName, "JobTitle") &&
			!strings.HasPrefix(fieldName, "BusinessUnit") &&
			!strings.HasPrefix(fieldName, "Gender") &&
			!strings.HasPrefix(fieldName, "Salary") &&
			!strings.HasPrefix(fieldName, "City") &&
			!strings.HasPrefix(fieldName, "Country") {
			message := "unknown field in new client"
			json.NewEncoder(w).Encode(map[string]string{"error": message})
			return
		}
	}

	// Handle errors in adding new client
	newClient.ID = strconv.Itoa(len(clients) + 2 + 2000)
	if err := addClient(newClient); err != nil {
		w.WriteHeader(http.StatusInternalServerError) // Set the status code to 500
		message := "Unable to create client"
		json.NewEncoder(w).Encode(map[string]string{"error": message})
		return
	}

	json.NewEncoder(w).Encode(newClient)
}
