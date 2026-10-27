package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"groupie-tracker/api"
)

// data contient toutes les données de l'API.
// Elle est déclarée hors de toute fonction : main la remplit, les handlers la lisent.
var data api.DataAPI

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

	// 3. Démarrer le serveur.
	fmt.Println("Le serveur est lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
