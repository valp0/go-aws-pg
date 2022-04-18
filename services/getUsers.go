package services

import (
	"github.com/valp0/go-aws-pg/repo"
)

func GetUsers() ([]repo.User, error) {
	return repo.GetUsers()
}
