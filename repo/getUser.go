package repo

import (
	_ "github.com/lib/pq"
)

// GetUser is the repository function that fetches a user, given its id.
func GetUser(id string) ([]User, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	user, err := getUser(db, id)
	if err != nil {
		return nil, err
	}

	return []User{user}, nil
}
