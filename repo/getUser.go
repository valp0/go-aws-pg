package repo

import (
	_ "github.com/lib/pq"
)

// GetUser is the repository function that fetches a user, given its id.
func (r repository) GetUser(id string) ([]User, error) {
	user, err := r.getUser(id)
	if err != nil {
		return nil, err
	}

	return []User{user}, nil
}
