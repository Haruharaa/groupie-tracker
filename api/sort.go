package api

import (
	"slices"
	"sort"
	"strings"
)

// SortArtists renvoie une copie de la liste, triée selon order :
// "name-asc" (A-Z), "name-desc" (Z-A) ou "creationDate" (du plus ancien au plus récent).
func SortArtists(artists []Artist, order string) []Artist {
	sorted := slices.Clone(artists) // une copie : on ne touche pas à la liste d'origine

	switch order {
	case "name-asc":
		sort.Slice(sorted, func(i, j int) bool {
			return strings.ToLower(sorted[i].Name) < strings.ToLower(sorted[j].Name)
		})
	case "name-desc":
		sort.Slice(sorted, func(i, j int) bool {
			return strings.ToLower(sorted[i].Name) > strings.ToLower(sorted[j].Name)
		})
	case "creationDate":
		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i].CreationDate < sorted[j].CreationDate
		})
	}

	return sorted
}
