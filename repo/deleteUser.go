package repo

import (
	_ "github.com/lib/pq"
)

// DeleteUser is the repository function to safely remove a user from the database.
func (r repository) DeleteUser(id string) error {
	if err := r.deleteFavsFromUser(id); err != nil {
		return err
	}

	if err := r.deleteUser(id); err != nil {
		return err
	}

	r.cleanFavs()
	return nil
}
