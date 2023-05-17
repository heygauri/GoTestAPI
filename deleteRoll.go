package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// Delete
func deleteClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	e := 0
	for i, item := range clients {
		if item.ID == params["id"] {
			clients = append(clients[:i], clients[i+1:]...)
			message := "Client deleted successfully"
			json.NewEncoder(w).Encode(map[string]string{"message": message})
			w.WriteHeader(http.StatusNoContent) // Set the status code to 204
			e = 1
			break
		}
	}
	// If no matching client is found, return a 404 error
	if e == 0 {
		http.Error(w, "Client not found", http.StatusNotFound)
	}
}
