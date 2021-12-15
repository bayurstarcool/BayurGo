package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/bayurstarcool/bayurGo/app/models"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You are on the about page.")
}

func (c *AppContext) IndexHandler(w http.ResponseWriter, r *http.Request) {
	// db := c.DB
	// user := []models.User{}
	// db.Find(&user)
	// json.NewEncoder(w).Encode(user)
	rnd.HTML(w, http.StatusOK, "welcome.html", nil)
}

func AuthHandler(next http.Handler) http.Handler {

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

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	user := context.Get(r, "user")
	// Maybe other operations on the database
	json.NewEncoder(w).Encode(user)
}

func (c *AppContext) TeaHandler(w http.ResponseWriter, r *http.Request) {
	params := context.Get(r, "params").(httprouter.Params)
	keyword := params.ByName("query")
	db := c.DB
	users := []models.User{}
	db.Where("email like ? ", "%"+keyword+"%").First(&users)
	if len(users) == 0 {
		json.NewEncoder(w).Encode("{users: Not Found}")
		return
	}
	// tea := getTea(c.db, params.ByName("id"))
	json.NewEncoder(w).Encode(users)
}
