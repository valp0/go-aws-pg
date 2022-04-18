package services

import (
	"github.com/valp0/go-aws-pg/repo"
)

func Ping() (string, error) {
	err := repo.Ping()
	if err != nil {
		return "", err
	}

	return "DB pinged sucessfully!", nil
}
