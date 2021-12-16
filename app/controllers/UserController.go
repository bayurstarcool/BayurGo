package controllers

import (
	"html/template"
	"net/http"

	"github.com/bayurstarcool/BayurGo/app/helpers"
	"github.com/bayurstarcool/BayurGo/app/models"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {
	tpls := []string{"views/layouts/backend.html", "views/backend/users/create.html"}
	rnd.Template(w, http.StatusOK, tpls, nil)
}
func (c *AppContext) UserStore(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	user := models.User{Name: name, Email: email, Password: &password}
	c.DB.Create(&user)
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}
func (c *AppContext) UserUpdate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	// password := r.FormValue("password")
	user := models.User{Name: name, Email: email}
	c.DB.Updates(&user)
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}
func (c *AppContext) UserEdit(w http.ResponseWriter, r *http.Request) {
	params := context.Get(r, "params").(httprouter.Params)
	id := params.ByName("id")
	db := c.DB
	user := models.User{}
	db.Where("id = ?", id).First(&user)
	tpls := []string{"views/layouts/backend.html", "views/backend/users/edit.html"}
	rnd.Template(w, http.StatusOK, tpls, user)
}
func (c *AppContext) UserIndex(w http.ResponseWriter, r *http.Request) {
	db := c.DB
	users := []models.User{}
	db.Find(&users)
	tpls := []string{"views/layouts/backend.html", "views/backend/users/index.html"}
	rnd.FuncMap(template.FuncMap{
		"inc": helpers.GetIncrement,
	})
	rnd.Template(w, http.StatusOK, tpls, users)
}
