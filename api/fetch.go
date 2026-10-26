package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchArtists() ([]Artist, error) {
	var artists []Artist

	err := fetchJSON("https://groupietrackers.herokuapp.com/api/artists", &artists) //&artists sait où ranger sans & il fait une copie et artists serait vide
	if err != nil {
		return nil, err
	}

	return artists, nil //On renvoit la boîte remplie
}

func FetchLocations() (LocationsIndex, error) {
	var locations LocationsIndex

	err := fetchJSON("https://groupietrackers.herokuapp.com/api/locations", &locations)
	if err != nil {
		return LocationsIndex{}, err
	}

	return locations, nil
}

func FetchDates() (DatesIndex, error) {
	var dates DatesIndex

	err := fetchJSON("https://groupietrackers.herokuapp.com/api/dates", &dates)
	if err != nil {
		return DatesIndex{}, err
	}

	return dates, nil
}

func FetchRelations() (RelationsIndex, error) {
	var relations RelationsIndex

	err := fetchJSON("https://groupietrackers.herokuapp.com/api/relation", &relations)
	if err != nil {
		return RelationsIndex{}, err
	}

	return relations, nil
}

// FetchJSON permet de récupérer des données JSON à partir d'une URL et de les décoder dans la bonne structure ( TARGET ).
func fetchJSON(url string, target any) error {
	resp, err := http.Get(url)
	if err != nil {
		return err //retourne l'erreur si ça échoue
	}
	defer resp.Body.Close() //Si rien trouvé on ferme la connexion
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("échec de la récupération des données: %s", resp.Status)
	}
	return json.NewDecoder(resp.Body).Decode(target) //il lit si tout va bien et le range avec target
}
