package postgres

import (
	"database/sql"
	"fmt"
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
func (b *storer) CreateDirectory(name string) (string, error) {
}

// AddEntry and return string "ok" or "fail"
func (b *storer) AddEntry(e *api.EntryRequest) (string, error) {

}

// SearchEntry
func (b *storer) SearchEntry(query string, limit uint) ([]*api.Entry, error) {

}

// Close handles any necessary cleanup
func (b *storer) Close() error {
	return b.db.Close()
}