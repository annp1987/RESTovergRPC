package postgres

import (
	"github.com/jinzhu/gorm"
)

type Directory struct {
	gorm.Model
	DirectoryName string `json:"directory_name"`
}

type Entry struct {
	gorm.Model
	DirectoryID uint
	Name string	`json:"name"`
	LastName string `json:"last_name"`
	PhNumber string `json:"ph_number"`
}