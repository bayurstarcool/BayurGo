package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/thedevsaddam/renderer"
)

var rnd *renderer.Render

type AppContext struct {
	DB *gorm.DB
}

func init() {
	print("\n\nWelcome, BayurGo Framework\n")
	// opts := renderer.Options{
	// 	ParseGlobPattern: "./views/*.html",
	// }
	print("BayurGo running in port 8080 or http://localhost:8080\n")
	// rnd = renderer.New(opts)
	rnd = renderer.New()
}
