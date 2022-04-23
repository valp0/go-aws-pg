package services

import (
	"encoding/json"
	"fmt"

	"github.com/valp0/go-aws-pg/repo"
)

// PostUser is the service function that will call the repository function to insert
// a user with the received data after validating the received user name and id.
func (s service) PostUser(decoder *json.Decoder) ([]repo.User, error) {
	var user repo.User
	err := decoder.Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("couldn't decode user data from the request body, %s", err.Error())
	}

	if user.ID == "" || user.Name == "" {
		return nil, fmt.Errorf("user_name and user_id are mandatory fields, they can't be empty")
	}

	if !validateUserId(user.ID) {
		return nil, fmt.Errorf("user_id can only have between 5 and 12 characters and can only contain letters, numbers and underscores")
	}

	if !validateUserName(user.Name) {
		return nil, fmt.Errorf("user_name can only have between 5 and 12 characters and can only contain letters, numbers and underscores")
	}

	return s.r.PostUser(user)
}
