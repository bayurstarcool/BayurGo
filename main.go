package main

import (
	"net/http"

	"github.com/bayurstarcool/BayurGo/app/controllers"
	"github.com/bayurstarcool/BayurGo/database"
	"github.com/bayurstarcool/BayurGo/route"
)

func main() {
	db := database.SetupDB()
	router := route.RouteApp(&controllers.AppContext{DB: db})
	http.ListenAndServe(":8080", router)
}
