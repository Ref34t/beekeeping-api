package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// getSenseBoxData handles GET requests for retrieving senseBox data
func getSenseBoxData(w http.ResponseWriter, r *http.Request) {
	// Extract senseBox ID from URL parameters
	vars := mux.Vars(r)
	id := vars["id"]

	// Simulate fetching data from openSenseMap (hardcoded for now)
	if id == "5eba5fbad46fb8001b799786" {
		senseBox := SenseBox{
			ID:   id,
			Temp: 22.5,
			Hum:  60.0,
		}

		// Set response header and encode data as JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(senseBox)
	} else {
		// Return 404 if ID is not found
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("SenseBox not found"))
	}
}
