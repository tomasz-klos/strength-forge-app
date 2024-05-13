package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rs/cors"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	mux := http.NewServeMux()

	corsHandler := cors.Default().Handler(mux)
	
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}