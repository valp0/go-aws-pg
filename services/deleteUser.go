package services

import (
	"fmt"

	"github.com/valp0/go-aws-pg/repo"
)

func DeleteUser(id string) ([]repo.User, error) {
	if !validateUserId(id) {
		return nil, fmt.Errorf("id can only have between 5 and 12 characters and can only contain letters, numbers and underscores")
	}

	if err := repo.DeleteUser(id); err != nil {
		return nil, err
	}

	return repo.GetUsers()
}
