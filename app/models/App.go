package models

import (
	"github.com/jinzhu/gorm"
)
type AppContext struct {
	DB *gorm.DB
}

func (c *AppContext) GetDB() (db *gorm.DB) {
	return c.DB
}
