package main

import (
	"net/http"

	"github.com/bayurstarcool/bayurGo/app/controllers"
	"github.com/bayurstarcool/bayurGo/database"
	"github.com/bayurstarcool/bayurGo/route"
)

func main() {
	db := database.SetupDB()
	router := route.RouteApp(&controllers.AppContext{DB: db})
	http.ListenAndServe(":8080", router)
}
