package repo

import (
	"database/sql"
	"fmt"
	"log"
)

// An auxiliary function to get all users from the users table.
func getUsers(db *sql.DB) ([]User, error) {
	users := []User{}
	sqlStatement := `SELECT * FROM Users;`

	rows, _ := db.Query(sqlStatement)
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.ThumbnailUrl)
		if err != nil {
			return nil, fmt.Errorf("an error occured while getting users, %v", err)
		}

		users = append(users, user)
	}

	return users, nil
}

// An auxiliary function to get a single user by its id.
func getUser(db *sql.DB, id string) (User, error) {
	var user User
	sqlStatement := `SELECT * FROM Users WHERE user_id = $1;`
	err := db.QueryRow(sqlStatement, id).Scan(&user.ID, &user.Name, &user.ThumbnailUrl)
	if err != nil {
		return User{}, fmt.Errorf("unable to fetch user with id %s, %v", id, err)
	}

	return user, nil
}

// An auxiliary repository function to insert a user into the users table.
func insertToUsers(db *sql.DB, user User) error {
	insertStatement := `INSERT INTO Users (user_id, user_name, user_thumbnail) VALUES ($1, $2, $3);`
	_, err := db.Exec(insertStatement, user.ID, user.Name, user.ThumbnailUrl)
	if err != nil {
		return fmt.Errorf("couldn't insert user with id %s, %v", user.ID, err)
	}

	log.Printf("User with id %s was inserted to users table successfully.\n", user.ID)
	return nil
}

// An auxiliary funciton to update a user name and/or thumbnail.
func updateUser(db *sql.DB, id string, user User) error {
	var updateStatement string
	var args []interface{}

	switch {
	case user.Name == "":
		updateStatement = `UPDATE Users SET user_thumbnail = $1 WHERE user_id = $2;`
		args = []interface{}{user.ThumbnailUrl, id}
	case user.ThumbnailUrl == "":
		updateStatement = `UPDATE Users SET user_name = $1 WHERE user_id = $2;`
		args = []interface{}{user.Name, id}
	default:
		updateStatement = `UPDATE Users SET user_name = $1, user_thumbnail = $2 WHERE user_id = $3;`
		args = []interface{}{user.Name, user.ThumbnailUrl, id}
	}

	_, err := db.Exec(updateStatement, args...)
	if err != nil {
		return fmt.Errorf("couldn't update user with id %s, %v", id, err)
	}

	log.Printf("User with id %s was inserted to users table successfully.\n", id)
	return nil
}

// Will delete a user from the Users table, given its id.
func deleteUser(db *sql.DB, id string) error {
	sqlStatement := `DELETE FROM Users WHERE user_id = $1;`
	_, err := db.Exec(sqlStatement, id)
	return err
}

// Checks a user is not already present in the table to avoid duplicity.
func alreadyInUsers(db *sql.DB, id string) bool {
	sqlStatement := `SELECT * FROM Users WHERE user_id = $1;`
	var user User

	err := db.QueryRow(sqlStatement, id).Scan(&user.ID, &user.Name, &user.ThumbnailUrl)
	return err == nil
}
