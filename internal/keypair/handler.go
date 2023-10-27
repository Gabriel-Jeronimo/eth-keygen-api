package keypair

import (
	"fmt"
	"log"
	"net/http"
)

func InitRoutes() {
	http.HandleFunc("/keypair", GenerateKeypairHandler)
	http.HandleFunc("/address", GetAddressHandler)
}

func GenerateKeypairHandler(w http.ResponseWriter, r *http.Request) {
	privateKey, publicKey, err := GenerateKeypair()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%q %q", privateKey, publicKey)
}

func GetAddressHandler(w http.ResponseWriter, r *http.Request) {
	address, err := GetAddress()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%q", address)
}
