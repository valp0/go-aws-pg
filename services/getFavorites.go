package services

import (
	"fmt"

	"github.com/valp0/go-aws-pg/repo"
)

// DeleteFavorite is the service function that will call the repository function
// to get a user's favorites after validating the given user id is valid.
func (s service) GetFavorites(id string) ([]repo.Favorite, error) {
	if !validateUserId(id) {
		return nil, fmt.Errorf("id can only have between 5 and 12 characters and can only contain letters, numbers and underscores")
	}

	return s.r.GetFavorites(id)
}
