// Building a REST ful API with Go
package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Client is model for client data
type Client struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	JobTitle     string `json:"job_title"`
	BusinessUnit string `json:"business_unit"`
	Gender       string `json:"gender"`
	Salary       string `json:"salary"`
	City         string `json:"city"`
	Country      string `json:"country"`
}

// Init clients var as a slice to store the database
var clients []Client

func main() {

	//generate mock data
	clients = append(clients,
		Client{ID: "2002",
			Name:         "Kai Le",
			JobTitle:     "Controls Engineer",
			BusinessUnit: "Manufacturing",
			Gender:       "Male",
			Salary:       "$92,368",
			City:         "Columbus",
			Country:      "United States"},
		Client{ID: "2003",
			Name:         "Robert Patel",
			JobTitle:     "Analyst",
			BusinessUnit: "Corporate",
			Gender:       "Male",
			Salary:       "$45,703",
			City:         "Chicago",
			Country:      "United States"},
		Client{ID: "2004",
			Name:         "Harper Castillo",
			JobTitle:     "Network Administrator",
			BusinessUnit: "Research & Development",
			Gender:       "Male",
			Salary:       "$83,576",
			City:         "Shanghai",
			Country:      "China"},
		Client{ID: "2005",
			Name:         "Jade Hu",
			JobTitle:     "Sr. Analyst",
			BusinessUnit: "Speciality Products",
			Gender:       "Female",
			Salary:       "$89,744",
			City:         "Chongqing",
			Country:      "China"},
		Client{ID: "2006",
			Name:         "Jose Wong",
			JobTitle:     "Director",
			BusinessUnit: "Manufacturing",
			Gender:       "Male",
			Salary:       "$150,558",
			City:         "Phoenix",
			Country:      "United States"},
		Client{ID: "2007",
			Name:         "Sophia Gutierrez`",
			JobTitle:     "Manager",
			BusinessUnit: "Special Products",
			Gender:       "Female",
			Salary:       "$102,649",
			City:         "Austin",
			Country:      "United States"},
		Client{ID: "2008",
			Name:         "Lillian Lewis",
			JobTitle:     "Technical Architect",
			BusinessUnit: "Research & Development",
			Gender:       "Female",
			Salary:       "$83,323",
			City:         "Seattle",
			Country:      "United States"})

	//initialize router
	router := mux.NewRouter()

	//endpoints
	router.HandleFunc("/client", getClients).Methods("GET")
	router.HandleFunc("/client/{id}", getClient).Methods("GET")
	router.HandleFunc("/client", createClient).Methods("POST")
	router.HandleFunc("/client/{id}", updateClient).Methods("POST")
	router.HandleFunc("/client/{id}", deleteClient).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))

}
