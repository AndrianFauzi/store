package main

import (
	"store/controllers"
	"store/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectionDB()
}

func main() {

	r := gin.Default()

	r.POST("/", controllers.PostCreate)
	r.GET("/", controllers.PostIndex)
	r.GET("/:id", controllers.PostShow)
	r.PUT("/:id", controllers.PostUpdate)
	r.DELETE("/:id", controllers.PostDelete)
	r.Run() 
}