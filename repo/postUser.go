package repo

import (
	"fmt"

	_ "github.com/lib/pq"
)

// PostUser is the repository function to insert new users to the users table.
func (r repository) PostUser(user User) ([]User, error) {
	err := r.insertToUsers(user)
	if err != nil {
		return nil, err
	}

	id := user.ID
	user = User{}
	user, err = r.getUser(id)
	if err != nil {
		return nil, fmt.Errorf("%s, verify user was added to table", err.Error())
	}

	return []User{user}, nil
}
