package repo

import (
	_ "github.com/lib/pq"
)

// GetFavorites is the repository function that fetches the favorites of a user, given their id.
func GetFavorites(id string) ([]Favorite, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return getFavorites(db, id)
}
