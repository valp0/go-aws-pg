package repo

import (
	"fmt"

	_ "github.com/lib/pq"
)

func UpdateUser(id string, user User) ([]User, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if !alreadyInUsers(db, id) {
		return nil, fmt.Errorf("id %s is not present in the users table", id)
	}

	if user.ID != "" {
		return nil, fmt.Errorf("user id cannot be modified")
	}

	if err = updateUser(db, id, user); err != nil {
		return nil, err
	}

	user = User{}
	user, err = getUser(db, id)
	if err != nil {
		return nil, err
	}

	return []User{user}, nil
}
