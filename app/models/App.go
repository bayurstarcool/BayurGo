package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Name  string
	Email string
}
type AppContext struct {
	DB *gorm.DB
}

func (c *AppContext) GetDB() (db *gorm.DB) {
	return c.DB
}
