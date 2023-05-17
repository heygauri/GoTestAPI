package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// Show
func getClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range clients {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// If no matching client is found, return a 404 error
	http.Error(w, "Client not found", http.StatusNotFound)
}
