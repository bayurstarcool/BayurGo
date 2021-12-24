package controllers

import (
	"net/http"

	"github.com/bayurstarcool/BayurGo/app/helpers"
)

func MyNotFound(w http.ResponseWriter, r *http.Request) {
	idn := struct {
		Name string
		Ver  string
	}{helpers.GetEnv("APP_NAME"), helpers.GetEnv("APP_VERSION")}
	tpls := []string{"views/layouts/app.html", "views/layouts/partial.html", "views/error404.html"}
	rnd.Template(w, r, http.StatusNotFound, tpls, idn)
}
