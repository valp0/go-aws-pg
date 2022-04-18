package repo

import (
	_ "github.com/lib/pq"
)

// GetUsers is the repository function to get all users from the users table.
func GetUsers() ([]User, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return getUsers(db)
}
