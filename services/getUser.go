package services

import (
	"fmt"

	"github.com/valp0/go-aws-pg/repo"
)

func GetUser(id string) ([]repo.User, error) {
	if !validateUserId(id) {
		return nil, fmt.Errorf("id can only have between 5 and 12 characters and can only contain letters, numbers and underscores")
	}

	return repo.GetUser(id)
}
