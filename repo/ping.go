package repo

import (
	_ "github.com/lib/pq"
)

// Ping is the repo function that will issue a ping to the database and return an error if not successful.
func Ping() error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Ping()
}
