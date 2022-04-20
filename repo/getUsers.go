package repo

import (
	_ "github.com/lib/pq"
)

// GetUsers is the repository function to get all users from the users table.
func (r repository) GetUsers() ([]User, error) {

	return r.getUsers()
}
