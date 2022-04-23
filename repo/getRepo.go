package repo

import "database/sql"

type repository struct {
	db *sql.DB
}

var repo repository

// Returns a local repository to handle a database.
func GetRepo() (repository, error) {
	if repo.db == nil {
		db, err := connectDB()
		repo.db = db
		if err != nil {
			return repository{}, err
		}
	}

	return repo, nil
}
