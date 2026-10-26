package api

// FindArtist cherche un artiste par son id si oui bool true sinon false retourne vide
func FindArtist(artists []Artist, id int) (Artist, bool) {
	for _, artist := range artists {
		if artist.ID == id {
			return artist, true
		}
	}
	return Artist{}, false
}

// FindRelation cherche les relations (lieux et dates) d'un artiste par son id.
func FindRelation(relations RelationsIndex, id int) (Relations, bool) {
	for _, relation := range relations.Index {
		if relation.ID == id {
			return relation, true
		}
	}
	return Relations{}, false
}