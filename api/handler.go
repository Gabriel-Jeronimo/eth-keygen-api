package api

import (
	"fmt"
	"net/http"
)

func GenerateKeypair(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Vamos por mais")
}

func GetAddress(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Lets go")
}
