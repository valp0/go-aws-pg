package repo

import (
	"fmt"

	_ "github.com/lib/pq"
)

// PostUser is the repository function to insert new users to the users table.
func PostUser(user User) ([]User, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = insertToUsers(db, user)
	if err != nil {
		return nil, err
	}

	id := user.ID
	user = User{}
	user, err = getUser(db, id)
	if err != nil {
		return nil, fmt.Errorf("%s, verify user was added to table", err.Error())
	}

	return []User{user}, nil
}
