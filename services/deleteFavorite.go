package services

import (
	"fmt"

	"github.com/valp0/go-aws-pg/repo"
)

// DeleteFavorite is the service function that will call the repository function to
// delete a favorite after validating both the given user id and video id are valid.
func (s service) DeleteFavorite(userId, vidId string) ([]repo.Favorite, error) {
	if !validateUserId(userId) {
		return nil, fmt.Errorf("id can only have between 5 and 12 characters and can only contain letters, numbers and underscores")
	}

	if !validateVidId(vidId) {
		return nil, fmt.Errorf("id can only contain letters, numbers, dashes and underscores, and must be 11 characters long")
	}

	if err := s.r.DeleteFavorite(userId, vidId); err != nil {
		return nil, err
	}

	return s.r.GetFavorites(userId)
}
