package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handlePublicacoes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	response := map[string]interface{}{
		"message":     "Publicações",
		"publicações": []string{"publicação 1", "publicação 2"},
	}

	json.NewEncoder(w).Encode(response)
}

func main() {

	r := mux.NewRouter()
	fmt.Println("Iniciando o servidor...")

	r.HandleFunc("/publicacoes", handlePublicacoes).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))

}
