package models

import (
	"github.com/jinzhu/gorm"
)

type AppContext struct {
	DB *gorm.DB
}
type M map[string]interface{}

func (c *AppContext) GetDB() (db *gorm.DB) {
	return c.DB
}
