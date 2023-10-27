package main

import (
	"log"
	"net/http"

	"github.com/Gabriel-Jeronimo/eth-keygen-api/api"
)

func main() {
	http.HandleFunc("/keypair", api.GenerateKeypair)
	http.HandleFunc("/address", api.GetAddress)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
