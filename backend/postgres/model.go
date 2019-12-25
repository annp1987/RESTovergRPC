package postgres

import (
	"github.com/jinzhu/gorm"
)

type Directory struct {
	gorm.Model
	DirectoryName string `gorm:"not null;unique"`
	Entries []Entry `gorm:"foreignkey:DirectoryRefer"`
}

type Entry struct {
	gorm.Model
	DirectoryRefer string
	Name string
	LastName string
	PhNumber string
}