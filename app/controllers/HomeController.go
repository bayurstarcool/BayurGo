package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/bayurstarcool/BayurGo/app/models"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	tpls := []string{"views/layouts/backend.html", "views/layouts/partial.html", "views/dashboard.html"}
	rnd.Template(w, r, http.StatusOK, tpls, nil)
}

func (c *AppContext) IndexHandler(w http.ResponseWriter, r *http.Request) {
	// db := c.DB
	// user := []models.User{}
	// db.Find(&user)
	// json.NewEncoder(w).Encode(user)
	tpls := []string{"views/layouts/app.html", "views/layouts/partial.html", "views/welcome.html"}
	rnd.Template(w, r, http.StatusOK, tpls, nil)
}

func (c *AppContext) AdminHandler(w http.ResponseWriter, r *http.Request) {
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
	json.NewEncoder(w).Encode(users)
}
