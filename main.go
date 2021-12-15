package main

import (
	"net/http"

	"github.com/bayurstarcool/bayurGo/app/controllers"
	"github.com/bayurstarcool/bayurGo/database"
	"github.com/bayurstarcool/bayurGo/route"
)

func main() {
	// db := sql.Open("postgres", "...")
	print("Welcome, BayurGo Framework\n")
	db := database.SetupDB()
	router := route.RouteApp(&controllers.AppContext{DB: db})
	http.ListenAndServe(":8080", router)
}
