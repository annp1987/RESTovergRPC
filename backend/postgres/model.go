package postgres

import (
	"github.com/jinzhu/gorm"
)

type Directory struct {
	gorm.Model
	DirectoryName string `gorm:"not null;unique" json:"directory_name"`
	Entries []Entry `gorm:"foreignkey:DirectoryRefer"`
}

type Entry struct {
	gorm.Model
	DirectoryRefer string `json:"directory_name"`
	Name string	`json:"name"`
	LastName string `json:"last_name"`
	PhNumber string	`json:"ph_number"`
}