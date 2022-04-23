package services

import (
	"encoding/json"
	"fmt"

	"github.com/valp0/go-aws-pg/repo"
)

// UpdateUser is the service function that will call the repository function to update
// the user with received id using received data after validating user id and content are valid.
func (s service) UpdateUser(id string, decoder *json.Decoder) ([]repo.User, error) {
	var user repo.User

	if !validateUserId(id) {
		return nil, fmt.Errorf("id can only have between 5 and 12 characters and can only contain letters, numbers and underscores")
	}

	err := decoder.Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("couldn't decode user data from the request body, %s", err.Error())
	}

	if user.Name == "" && user.ThumbnailUrl == "" && user.ID == "" {
		return nil, fmt.Errorf("couldn't decode user data from the request body, body is empty")
	}

	return s.r.UpdateUser(id, user)
}
