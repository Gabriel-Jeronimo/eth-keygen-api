package keypair

import (
	"encoding/json"
	"net/http"
)

type GenerateKeypairHandlerResponse struct {
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
}

type GetAddressHandlerResponse struct {
	Address string `json:"address"`
}

func InitRoutes() {
	http.HandleFunc("/keypair", GenerateKeypairHandler)
	http.HandleFunc("/address", GetAddressHandler)
}

func GenerateKeypairHandler(w http.ResponseWriter, r *http.Request) {
	privateKey, publicKey, err := GenerateKeypair()

	if err != nil {
		http.Error(w, "Failed to generate a new keypair: "+err.Error(), http.StatusInternalServerError)
		return
	}

	responseData := GenerateKeypairHandlerResponse{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}

	jsonResponse, err := json.Marshal(responseData)

	if err != nil {
		http.Error(w, "Failed to create JSON response"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}

func GetAddressHandler(w http.ResponseWriter, r *http.Request) {
	publicKey := r.URL.Query().Get("publicKey")

	if publicKey == "" {
		http.Error(w, "Parameter 'publicKey' not provided in the query string", http.StatusBadRequest)
		return
	}

	address, err := GetAddress(publicKey)

	if err != nil {
		http.Error(w, "Failed to retrieve the address from the public key: "+err.Error(), http.StatusInternalServerError)
		return
	}

	responseData := GetAddressHandlerResponse{
		Address: address,
	}

	jsonResponse, err := json.Marshal(responseData)

	if err != nil {
		http.Error(w, "Failed to create JSON response"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
