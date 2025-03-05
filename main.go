package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const appVersion = "v0.0.1"

func main() {
	versionFlag := flag.Bool("version", false, "Print the current app version")
	flag.Parse()

	if *versionFlag {
		fmt.Printf("App Version: %s\n", appVersion)
		os.Exit(0)
	}

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/sensebox/{id}", getSenseBoxData).Methods("GET")

	log.Println("Server started on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
