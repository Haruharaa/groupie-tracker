package main

import (
	"fmt"
	"groupie-tracker/api"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// data contient toutes les données de l'API.
// Elle est déclarée hors de toute fonction : main la remplit, les handlers la lisent.
var data api.DataAPI

// ArtistPage regroupe artist et relation
type ArtistPage struct {
	Artist   api.Artist
	Relation api.Relations
}

// accueil affiche la page d'accueil avec la liste de tous les artistes.
func accueil(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Cette page n'existe pas", http.StatusNotFound)
		return
	}
	// 1. Charger le fichier HTML du template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		// Fichier absent ou faute de syntaxe dans le HTML : erreur 500 pour le visiteur,
		// mais le serveur continue de tourner pour les autres
		log.Println("erreur de chargement du template :", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	// 2. Remplir le template avec la liste des artistes et l'envoyer au navigateur.
	// Dans index.html, le {{ range . }} parcourt cette liste.
	err = tmpl.Execute(w, data.Artists)
	if err != nil {
		log.Println("erreur d'affichage du template :", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}
}

// pageArtiste affiche la page détail d'un artiste : /artist?id=3
func pageArtiste(w http.ResponseWriter, r *http.Request) {
	// 1. Lire l'id dans l'adresse
	idTexte := r.URL.Query().Get("id")

	// 2. Le convertir en nombre
	id, err := strconv.Atoi(idTexte)
	if err != nil {
		http.Error(w, "Pas un nombre", http.StatusBadRequest)
		return
	}

	// 3. Trouver l'artiste
	artist, ok := api.FindArtist(data.Artists, id)
	if !ok {
		http.Error(w, "N'existe pas", http.StatusNotFound)
		return
	}
	//trouver les concerts
	relation, ok := api.FindRelation(data.Relations, id)
	if !ok {
		http.Error(w, "Concerts introuvables", http.StatusInternalServerError)
		return
	}
	//5. Ranger l'artiste + les concerts ensemble
	page := ArtistPage{Artist: artist, Relation: relation}
	// 6. Pour l'instant, on affiche en texte le nom et le nombre de lieux
	fmt.Fprintf(w, "%s : %d lieux de concerts\n", page.Artist.Name, len(page.Relation.DatesLocations))
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
	// Les fichiers du dossier statics (CSS, images) sont servis à l'adresse /statics/
	http.Handle("/statics/", http.StripPrefix("/statics/", http.FileServer(http.Dir("statics"))))
	http.HandleFunc("/", accueil)
	http.HandleFunc("/artist", pageArtiste)

	// 3. Démarrer le serveur.
	fmt.Println("Le serveur est lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
