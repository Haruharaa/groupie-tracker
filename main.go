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

	// ===== TESTS TEMPORAIRES (à supprimer quand tout marche) =====

	// Test de CleanDate : les deux lignes doivent afficher 23-08-2019
	fmt.Println("--- Test CleanDate ---")
	fmt.Println(api.CleanDate("*23-08-2019")) // avec astérisque
	fmt.Println(api.CleanDate("23-08-2019"))  // sans astérisque : rien ne doit changer

	// Test de FormatLocation avec des exemples écrits à la main
	fmt.Println("--- Test FormatLocation ---")
	fmt.Println(api.FormatLocation("north_carolina-usa"))    // attendu : North Carolina, Usa
	fmt.Println(api.FormatLocation("paris-france"))          // attendu : Paris, France
	fmt.Println(api.FormatLocation("rio_de_janeiro-brazil")) // attendu : Rio De Janeiro, Brazil
	fmt.Println(api.FormatLocation("bizarre"))               // attendu : bizarre (pas de "-", rien ne change)
	fmt.Println(api.FormatLocation(""))                      // attendu : une ligne vide, sans plantage

	// Test sur les vraies données : les concerts de Queen (id 1)
	fmt.Println("--- Concerts de Queen ---")
	relation, ok := api.FindRelation(data.Relations, 1)
	if ok {
		for lieu, dates := range relation.DatesLocations {
			fmt.Println(api.FormatLocation(lieu)) // le lieu, rendu lisible
			for _, date := range dates {
				fmt.Println("   ", api.CleanDate(date)) // chaque date, sans astérisque
			}
		}
	}

	// ===== FIN DES TESTS =====

	// 2. Relier les routes aux handlers.
	http.HandleFunc("/", accueil)

	// 3. Démarrer le serveur.
	fmt.Println("Le serveur est lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
