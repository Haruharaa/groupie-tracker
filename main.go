package main

import (
	"fmt"
	"groupie-tracker/api"
	"log"
	"net/http"
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
	locations, err := api.FetchLocations()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Lieux récupérés avec succès :")
	for _, loc := range locations.Index {
		fmt.Printf("ID: %d, %d lieux : %v\n", loc.ID, len(loc.Locations), loc.Locations)
	}
	http.HandleFunc("/", accueil)
	fmt.Println("Le serveur est lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
