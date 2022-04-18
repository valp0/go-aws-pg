package services

import (
	"fmt"

	"github.com/valp0/go-aws-pg/repo"
)

func GetFavorites(id string) ([]repo.Favorite, error) {
	if !validateUserId(id) {
		return nil, fmt.Errorf("id can only have between 5 and 12 characters and can only contain letters, numbers and underscores")
	}

	return repo.GetFavorites(id)
}
