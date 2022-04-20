package services

import (
	"encoding/json"
	"fmt"

	"github.com/valp0/go-aws-pg/repo"
)

// Will insert a user with the received data.
func (s service) PostFavorite(userId string, decoder *json.Decoder) ([]repo.Favorite, error) {
	var fav repo.Favorite
	err := decoder.Decode(&fav)
	if err != nil {
		return nil, fmt.Errorf("couldn't decode video data from the request body, %s", err.Error())
	}

	if !validateVidId(fav.ID) {
		return nil, fmt.Errorf("id can only contain letters, numbers, dashes and underscores, and must be 11 characters long")
	}

	if !validateVidTitle(fav.Title) {
		return nil, fmt.Errorf("video title can't be empty")
	}

	return s.r.PostFavorite(userId, fav)
}
