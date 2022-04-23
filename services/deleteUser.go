package services

import (
	"fmt"

	"github.com/valp0/go-aws-pg/repo"
)

// DeleteUser is the service function that will call the repository
// function to delete a user after validating the given user id is valid.
func (s service) DeleteUser(id string) ([]repo.User, error) {
	if !validateUserId(id) {
		return nil, fmt.Errorf("id can only have between 5 and 12 characters and can only contain letters, numbers and underscores")
	}

	if err := s.r.DeleteUser(id); err != nil {
		return nil, err
	}

	return s.r.GetUsers()
}
