package repo

import (
	"fmt"

	_ "github.com/lib/pq"
)

// DeleteFavorite is the repository function to delete a favorite.
func (r repository) DeleteFavorite(userId, favId string) error {
	if user, err := r.getUser(userId); user.ID == "" {
		if err != nil {
			return err
		}
		return fmt.Errorf("user with id %s does not exist", userId)
	}

	if err := r.deleteUserFav(userId, favId); err != nil {
		return err
	}

	return nil
}
