package keypair

import (
	"fmt"
	"net/http"
)

func InitRoutes() {
	http.HandleFunc("/keypair", GenerateKeypairHandler)
	http.HandleFunc("/address", GetAddressHandler)
}

// TODO: Create the return structure
func GenerateKeypairHandler(w http.ResponseWriter, r *http.Request) {
	privateKey, publicKey, err := GenerateKeypair()

	if err != nil {
		http.Error(w, "Failed to generate a new keypair: "+err.Error(), http.StatusInternalServerError)
	}

	fmt.Printf("%q %q", privateKey, publicKey)
}

func GetAddressHandler(w http.ResponseWriter, r *http.Request) {
	address, err := GetAddress("")

	if err != nil {
		http.Error(w, "Failed to retrieve the address from the public key: "+err.Error(), http.StatusInternalServerError)
	}

	fmt.Printf("%q", address)
}
