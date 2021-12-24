package controllers

import (
	"html/template"
	"net/http"

	"github.com/bayurstarcool/BayurGo/app/helpers"
	"github.com/bayurstarcool/BayurGo/app/models"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	tpls := []string{"views/layouts/backend.html", "views/backend/todos/create.html"}
	rnd.Template(w, r, http.StatusOK, tpls, nil)
}
func (c *AppContext) TodoStore(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := r.FormValue("name")
	note := r.FormValue("note")
	todo := models.Todo{Name: name, Note: note}
	c.DB.Create(&todo)
	http.Redirect(w, r, "/todos", http.StatusSeeOther)
}
func (c *AppContext) TodoUpdate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := r.FormValue("name")
	note := r.FormValue("note")
	// password := r.FormValue("password")
	todo := models.Todo{Name: name, Note: note}
	c.DB.Updates(&todo)
	http.Redirect(w, r, "/todos", http.StatusSeeOther)
}
func (c *AppContext) TodoEdit(w http.ResponseWriter, r *http.Request) {
	params := context.Get(r, "params").(httprouter.Params)
	id := params.ByName("id")
	db := c.DB
	todo := models.Todo{}
	db.Where("id = ?", id).First(&todo)
	tpls := []string{"views/layouts/backend.html", "views/backend/todos/edit.html"}
	rnd.Template(w, r, http.StatusOK, tpls, todo)
}
func (c *AppContext) TodoIndex(w http.ResponseWriter, r *http.Request) {
	db := c.DB
	todos := []models.Todo{}
	db.Find(&todos)
	tpls := []string{"views/layouts/backend.html", "views/backend/todos/index.html"}
	rnd.FuncMap(template.FuncMap{
		"inc": helpers.GetIncrement,
	})
	rnd.Template(w, r, http.StatusOK, tpls, todos)
}
