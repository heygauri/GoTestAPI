package main

import (
	"encoding/json"
	"net/http"
)

// Index
//func getClients(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//
//	// Handle errors in getting clients
//	if err := json.NewEncoder(w).Encode(clients); err != nil {
//		w.WriteHeader(http.StatusInternalServerError) // Set the status code to 500
//		message := "Unable to get clients"
//		json.NewEncoder(w).Encode(map[string]string{"error": message})
//		return
//	}
//}

func getClients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Handle errors in getting clients
	err := json.NewEncoder(w).Encode(clients)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // Set the status code to 500
		message := "Unable to get clients"
		errorRes := map[string]string{"error": message}
		jsonErr := json.NewEncoder(w).Encode(errorRes)
		if jsonErr != nil {
			http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
}
