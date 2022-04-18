package services

import (
	"fmt"

	"github.com/valp0/go-aws-pg/repo"
)

func DeleteFavorite(userId, vidId string) ([]repo.Favorite, error) {
	if !validateUserId(userId) {
		return nil, fmt.Errorf("id can only have between 5 and 12 characters and can only contain letters, numbers and underscores")
	}

	if !validateVidId(vidId) {
		return nil, fmt.Errorf("id can only contain letters, numbers, dashes and underscores, and must be 11 characters long")
	}

	return repo.DeleteFavorite(userId, vidId)
}
