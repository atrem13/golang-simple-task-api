package routes

import (
	"try-simple-api-task/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/tasks", controllers.FindTasks)
	r.POST("/tasks", controllers.CreateTasks)
	r.GET("/tasks/:id", controllers.FindTasks)
	r.PATCH("/tasks/:id", controllers.UpdateTasks)
	r.DELETE("/tasks/:id", controllers.DeleteTasks)
	return r
}
