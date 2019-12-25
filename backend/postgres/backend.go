package postgres

import (
	"github.com/jinzhu/gorm"

	// postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Directory struct {
	gorm.Model
	DirectoryName string
}

type Entry struct {
	gorm.Model
	DirectoryID uint
	Name string
	LastName string
	PhNumber string
}