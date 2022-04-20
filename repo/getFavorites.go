package repo

import (
	_ "github.com/lib/pq"
)

// GetFavorites is the repository function that fetches the favorites of a user, given their id.
func (r repository) GetFavorites(id string) ([]Favorite, error) {

	return r.getFavorites(id)
}
