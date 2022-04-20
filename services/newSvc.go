package services

import "github.com/valp0/go-aws-pg/repo"

type service struct {
	r servicer
}

type servicer interface {
	getter
	poster
	deleter
	updater
	pinger
}

type getter interface {
	GetUsers() ([]repo.User, error)
	GetFavorites(string) ([]repo.Favorite, error)
	GetUser(string) ([]repo.User, error)
}

type deleter interface {
	DeleteFavorite(string, string) error
	DeleteUser(string) error
}

type updater interface {
	UpdateUser(string, repo.User) ([]repo.User, error)
}

type poster interface {
	PostFavorite(string, repo.Favorite) ([]repo.Favorite, error)
	PostUser(repo.User) ([]repo.User, error)
}

type pinger interface {
	Ping() error
}

func NewService(r servicer) service {
	return service{r}
}
