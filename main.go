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
	r.Run() 
}