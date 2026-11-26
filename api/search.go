package api

import "strings"

// SearchArtists renvoie les artistes dont le nom contient query,
// sans tenir compte des majuscules.
func SearchArtists(artists []Artist, query string) []Artist {
	query = strings.ToLower(query)

	var results []Artist // la liste des artistes trouvés, vide au départ
	for _, artist := range artists {
		name := strings.ToLower(artist.Name) // le nom de l'artiste en minuscules
		if strings.Contains(name, query) {
			results = append(results, artist) // on ajoute l'artiste aux résultats
		}
	}
	return results
}
