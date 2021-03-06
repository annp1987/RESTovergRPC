package server

import (
	"github.com/annp1987/RESTovergRPC/backend"
	api "github.com/annp1987/RESTovergRPC/directory"
	"golang.org/x/net/context"
)

// Directory implements the DirectoryServer
type Directory struct {
	backend backend.Backend
}

func NewDirectoryServer(dbUrl map[string]string) (*Directory, error) {

	db, err := get_backend(dbUrl)

	if err != nil {
		return nil, err
	}
	return &Directory{backend: db}, err
}

// CreateDirectory create a directory to stores entries
func (d *Directory) CreateDirectory(ctx context.Context, req*api.DirectoryRequest) (*api.SuccessResponse, error) {

	success, err := d.backend.CreateDirectory(req.DirectoryName)

	return &api.SuccessResponse{Success: success}, err
}

// AddEntry creates a new entry
func (d *Directory) AddEntry(ctx context.Context, req*api.EntryRequest) (*api.SuccessResponse, error) {

	success, err := d.backend.AddEntry(req)

	return &api.SuccessResponse{Success: success}, err
}

// SearchEntity finds existing entities matching a query
func (d *Directory) SearchEntry(ctx context.Context, req*api.SearchEntryRequest) (*api.SearchEntriesResponse, error) {

	result, err := d.backend.SearchEntry(req.Query)

	if err != nil {
		return nil, err
	}
	resp := &api.SearchEntriesResponse{Entries: result}

	return resp, nil
}

// Cleanup
func (d *Directory) Close() error {
	return  d.backend.Close()
}