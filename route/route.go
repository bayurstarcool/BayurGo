package route

import (
	"net/http"

	"github.com/bayurstarcool/BayurGo/app/controllers"
	"github.com/gorilla/context"
	"github.com/justinas/alice"
)

func RouteApp(appContext *controllers.AppContext) (r *router) {
	appC := appContext
	commonHandlers := alice.New(context.ClearHandler, loggingHandler, recoverHandler)
	router := NewRouter()
	router.NotFound = http.HandlerFunc(controllers.MyNotFound)
	router.ServeFiles("/assets/*filepath", http.Dir("assets"))
	router.Get("/admin", commonHandlers.Append(appC.AuthHandler).ThenFunc(appC.AdminHandler))
	router.Get("/dashboard", commonHandlers.ThenFunc(controllers.Dashboard))

	//User resources
	router.Get("/users", commonHandlers.ThenFunc(appC.UserIndex))
	router.POST("/users", appC.UserStore)
	// router.POST("/token", controllers.TokenStore)
	router.POST("/users/:id", appC.UserUpdate)
	router.Get("/new/user", commonHandlers.ThenFunc(controllers.UserCreate))
	router.Get("/users/:id/edit", commonHandlers.ThenFunc(appC.UserEdit))

	//Todos resources
	router.Get("/todos", commonHandlers.ThenFunc(appC.TodoIndex))
	router.POST("/todos", appC.TodoStore)
	router.POST("/todos/:id", appC.TodoUpdate)
	router.Get("/new/todo", commonHandlers.ThenFunc(controllers.TodoCreate))
	router.Get("/todos/:id/edit", commonHandlers.ThenFunc(appC.TodoEdit))

	router.Get("/", commonHandlers.ThenFunc(appC.IndexHandler))
	router.Get("/teas/:query", commonHandlers.ThenFunc(appC.TeaHandler))
	return router
}
