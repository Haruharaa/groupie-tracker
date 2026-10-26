package main

import (
	"net/http"
	"fmt"
	"log"
)

func accueil (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/",accueil)
	fmt.Println("Le serveur est lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}