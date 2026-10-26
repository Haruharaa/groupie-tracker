package main

import (
	"fmt"
	"log"
	"net/http"
	"groupie-tracker/api"
)

func accueil(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	artists, err := api.FetchArtists()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Artistes récupérés avec succès:")
	for _, artist := range artists {
		fmt.Printf("ID: %d, Nom: %s\n", artist.ID, artist.Name)
	}
	http.HandleFunc("/", accueil)
	fmt.Println("Le serveur est lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
