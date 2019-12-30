package server

import (
	"errors"
	storage "github.com/annp1987/RESTovergRPC/backend/postgres"
	"github.com/annp1987/RESTovergRPC/backend"
)

var (
	InvalidDBType = errors.New("invalid db type")
)

// Config backend include Type and db url
func get_backend(dbUrl map[string]string) (backend.Backend, error) {
	switch dbUrl["Type"] {
	case "postgres":
		return storage.New(dbUrl), nil
	default:
		return nil, InvalidDBType
	}
}