package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func getSenseBoxData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "5eba5fbad46fb8001b799786" {
		senseBox := SenseBox{
			ID:   id,
			Temp: 22.5,
			Hum:  60.0,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(senseBox)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("SenseBox not found"))
	}
}
