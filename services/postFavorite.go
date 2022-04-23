package services

import (
	"encoding/json"
	"fmt"

	"github.com/valp0/go-aws-pg/repo"
)

// PostFavorite is the service function that will call the repository function to insert
// a favorite with the received data after checking the video id and title are valid.
func (s service) PostFavorite(userId string, decoder *json.Decoder) ([]repo.Favorite, error) {
	var fav repo.Favorite
	err := decoder.Decode(&fav)
	if err != nil {
		return nil, fmt.Errorf("couldn't decode video data from the request body, %s", err.Error())
	}

	if !validateUserId(userId) {
		return nil, fmt.Errorf("user_id can only have between 5 and 12 characters and can only contain letters, numbers and underscores")
	}

	if !validateVidId(fav.ID) {
		return nil, fmt.Errorf("id can only contain letters, numbers, dashes and underscores, and must be 11 characters long")
	}

	if !validateVidTitle(fav.Title) {
		return nil, fmt.Errorf("video title can't be empty")
	}

	return s.r.PostFavorite(userId, fav)
}
