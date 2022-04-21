package handlers

import (
	"encoding/json"

	"github.com/valp0/go-aws-pg/repo"
	"github.com/valp0/go-aws-pg/services"
)

type handler struct {
	s svcHandler
}

type svcHandler interface {
	getter
	poster
	deleter
	updater
	pinger
	tokener
}

type getter interface {
	GetFavorites(string) ([]repo.Favorite, error)
	GetUser(string) ([]repo.User, error)
	GetUsers() ([]repo.User, error)
}

type deleter interface {
	DeleteFavorite(string, string) ([]repo.Favorite, error)
	DeleteUser(string) ([]repo.User, error)
}

type updater interface {
	UpdateUser(string, *json.Decoder) ([]repo.User, error)
}

type poster interface {
	PostFavorite(string, *json.Decoder) ([]repo.Favorite, error)
	PostUser(*json.Decoder) ([]repo.User, error)
}

type pinger interface {
	Ping() (string, error)
}

type tokener interface {
	GetToken(*json.Decoder) (*services.Client, error)
}

func NewHandler(s svcHandler) handler {
	return handler{s}
}
