package repo

import (
	_ "github.com/lib/pq"
)

// DeleteUser will perform the necessary actions in order to safely remove a user from the database.
func DeleteUser(id string) error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if err := deleteFavsFromUser(db, id); err != nil {
		return err
	}

	if err := deleteUser(db, id); err != nil {
		return err
	}

	cleanFavs(db)
	return nil
}
