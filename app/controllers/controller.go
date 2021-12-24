package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/bayurstarcool/BayurGo/vendors/thedevsaddam/renderer"
	"github.com/gorilla/context"
	"github.com/jinzhu/gorm"
)

var rnd *renderer.Render

type AppContext struct {
	DB *gorm.DB
}

func (c *AppContext) AuthHandler(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")
		user, err := map[string]interface{}{}, errors.New("test")
		// user, err := getUser(c.db, authToken)
		log.Println(authToken)

		if err != nil {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		context.Set(r, "user", user)
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
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
