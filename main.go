package main

import (
	"try-simple-api-task/models"
	"try-simple-api-task/routes"
)

func main() {
	db := models.SetupDB()
	db.AutoMigrate(&models.Task{})

	r := routes.SetupRoutes(db)
	r.Run()
}
