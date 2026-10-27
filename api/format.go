package api

import "strings"

func CleanDate(date string) string {
	return strings.TrimPrefix(date, "*") // enlève "*" au début
}

func FormatLocation(location string) string {
	location = strings.ReplaceAll(location, "_", " ") //remplace les _

	parts := strings.Split(location, "-") //sépare les mots à -

	if len(parts) != 2 {
		return location
	}
	city := capitalizeWords(parts[0])

	country := capitalizeWords(parts[1])
	if len(country) <= 3 {

	}
}

// capitalize met une majuscule au début de chaque mot (merci captain obvious)
func capitalize(word string) string { // Me permet de découper le mot pour mettre une majuscule au début
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(word[:1]) + word[1:]
}

func capitalizeWords(text string) string {
	words := strings.Fields(text)
	for i, word := range words { // parcours la liste de mot
		words[i] = capitalize(word) // remplace le mot par sa version avec capitalize
	}
	return strings.Join(words, " ") //Recolle les mots avec un espace
}
