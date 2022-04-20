package services

import (
	"github.com/valp0/go-aws-pg/repo"
)

func (s service) GetUsers() ([]repo.User, error) {
	return s.r.GetUsers()
}
