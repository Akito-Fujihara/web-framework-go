package main

import (
	"github.com/Akito-Fujihara/web-framework-go/controllers"
	"github.com/Akito-Fujihara/web-framework-go/framework"
)

func main() {
	engine := framework.NewEngine()
	router := engine.Router

	router.Get("/lists", controllers.ListsController)
	router.Get("/users", controllers.UsersController)
	router.Get("/students", controllers.StudentsController)
	engine.Run()
}
