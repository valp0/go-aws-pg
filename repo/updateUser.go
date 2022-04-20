package repo

import (
	"fmt"

	_ "github.com/lib/pq"
)

func (r repository) UpdateUser(id string, user User) ([]User, error) {
	if !r.alreadyInUsers(id) {
		return nil, fmt.Errorf("id %s is not present in the users table", id)
	}

	if user.ID != "" {
		return nil, fmt.Errorf("user id cannot be modified")
	}

	if err := r.updateUser(id, user); err != nil {
		return nil, err
	}

	user = User{}
	user, err := r.getUser(id)
	if err != nil {
		return nil, err
	}

	return []User{user}, nil
}
