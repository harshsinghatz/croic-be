package main

import (
	"croic/controllers"
	"croic/initializers"
	"croic/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.DBConnect()
	initializers.SyncDB()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	r.POST("/validate", middleware.UserAuth, controllers.Validate)
	r.POST("/todos", middleware.UserAuth, controllers.AddTodo)
	r.GET("/todos", middleware.UserAuth, controllers.GetAllTodos)

	r.POST("/todos/:id", middleware.UserAuth, controllers.UpdateTodoById)

	r.Run()
}
