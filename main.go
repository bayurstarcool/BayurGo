package main

import (
	"net/http"

	"github.com/bayurstarcool/bayurGo/app/models"
	"github.com/bayurstarcool/bayurGo/route"
)

func main() {
	// db := sql.Open("postgres", "...")
	print("Welcome, BayurGo Framework\n")
	router := route.RouteApp(&models.AppContext{})
	http.ListenAndServe(":8080", router)
}
