package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// fetchJSON télécharge le JSON d'une adresse et le range dans target.
func fetchJSON(url string, target any) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("échec de la récupération de %s : %s", url, resp.Status)
	}
	return json.NewDecoder(resp.Body).Decode(target)
}

func FetchArtists() ([]Artist, error) {
	var artists []Artist

	err := fetchJSON("https://groupietrackers.herokuapp.com/api/artists", &artists)
	if err != nil {
		return nil, err
	}

	return artists, nil
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

// LoadData appelle les quatre Fetch et range tout dans un seul classeur.
func LoadData() (DataAPI, error) {
	var data DataAPI
	var err error

	data.Artists, err = FetchArtists()
	if err != nil {
		return DataAPI{}, err
	}

	data.Locations, err = FetchLocations()
	if err != nil {
		return DataAPI{}, err
	}

	data.Dates, err = FetchDates()
	if err != nil {
		return DataAPI{}, err
	}

	data.Relations, err = FetchRelations()
	if err != nil {
		return DataAPI{}, err
	}

	return data, nil
}
