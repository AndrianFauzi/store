package main

import (
	"store/initializers"
	"store/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectionDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Product{})
}