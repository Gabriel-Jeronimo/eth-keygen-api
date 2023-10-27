package main

import (
	"log"
	"net/http"

	"github.com/Gabriel-Jeronimo/eth-keygen-api/internal/keypair"
)

func main() {
	keypair.InitRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
