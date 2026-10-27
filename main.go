package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie-tracker/api"
)

// data contient toutes les données de l'API.
// Elle est déclarée hors de toute fonction : main la remplit, les handlers la lisent.
var data api.DataAPI

func accueil(w http.ResponseWriter, r *http.Request) {
	if len(data.Artists) == 0 {
		fmt.Fprintln(w, "Aucun artiste n'a été chargé.")
		return
	}
	fmt.Fprintf(w, "%d artistes chargés. Le premier est %s.\n", len(data.Artists), data.Artists[0].Name)
}

func main() {
	// 1. Charger toutes les données de l'API dans la variable globale.
	var err error
	data, err = api.LoadData()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(data.Artists), "artistes chargés")
	
	// 2. Relier les routes aux handlers.
	http.HandleFunc("/", accueil)

	// 3. Démarrer le serveur.
	fmt.Println("Le serveur est lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
