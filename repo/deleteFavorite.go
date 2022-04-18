package repo

import (
	"fmt"

	_ "github.com/lib/pq"
)

// DeleteFavorite is the repository function to delete a favorite.
func DeleteFavorite(userId, favId string) ([]Favorite, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if user, err := getUser(db, userId); user.ID == "" {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("user with id %s does not exist", userId)
	}

	if err = deleteUserFav(db, userId, favId); err != nil {
		return nil, err
	}

	return GetFavorites(userId)
}
