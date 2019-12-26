package postgres

import (
	"database/sql"
	"fmt"
	api "github.com/annp1987/RESTovergRPC/directory"
	"github.com/annp1987/RESTovergRPC/backend"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	)

type storer struct {
	db *sql.DB
}

func New(dburl map[string]string) (backend.Backend, error) {
	//url would look like "host=myhost port=myport user=gorm dbname=gorm password=mypassword"
	url := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dburl["Host"], dburl["Port"], dburl["User"], dburl["Type"], dburl["Password"])
	db, err := sql.Open("postgres", url)

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Directory{}, &Entry{})
	return &storer{db:db}, nil
}

// CreateDirectory
func (s *storer) CreateDirectory(name string) (string, error) {
	s.db.Create(&Directory{DirectoryName: name})
	return "OK", nil
}

// AddEntry and return string "ok" or "fail"
func (s *storer) AddEntry(e *api.EntryRequest) (string, error) {
	s.db.Create(&Entry{DirectoryRefer: e.DirectoryName, Name: e.Name, LastName: e.LastName, PhNumber: e.PhNumber})
	return "OK", nil
}

// SearchEntry
func (s *storer) SearchEntry(query string) ([]*api.Entry, error) {
	var users []*api.Entry
	query_map := api.ExtractQuery(query)
	s.db.Where(query_map).Find(users)
	return users, nil
}

// Close handles any necessary cleanup
func (s *storer) Close() error {
	return s.db.Close()
}