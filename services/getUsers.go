package services

import (
	"github.com/valp0/go-aws-pg/repo"
)

// GetUser is the service function that will call the repository function to get all users.
func (s service) GetUsers() ([]repo.User, error) {
	return s.r.GetUsers()
}
