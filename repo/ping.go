package repo

import (
	_ "github.com/lib/pq"
)

// Ping is the repo function that will issue a ping to the database and return an error if not successful.
func (r repository) Ping() error {

	return r.db.Ping()
}
