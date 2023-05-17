package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// Update
func updateClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range clients {
		if item.ID == params["id"] {
			clients = append(clients[:i], clients[i+1:]...)
			var newClient Client
			err := json.NewDecoder(r.Body).Decode(&newClient)
			if err != nil {
				// If the request body is not valid JSON, return a 400 error
				http.Error(w, "Invalid request body", http.StatusBadRequest)
				return
			}
			newClient.ID = params["id"]
			clients = append(clients, newClient)
			json.NewEncoder(w).Encode(newClient)
			return
		}
	}
	// If no matching client is found, return a 404 error
	http.Error(w, "Client not found", http.StatusNotFound)
}
