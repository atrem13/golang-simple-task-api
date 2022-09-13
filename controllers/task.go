package controllers

import (
	"net/http"
	"time"
	"try-simple-api-task/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateTaskInput struct {
	AssingedTo string `json:"assignedTo"`
	Task       string `json:"task"`
	Deadline   string `json:"deadline`
}

type UpdateTaskInput struct {
	AssingedTo string `json:"assignedTo"`
	Task       string `json:"task"`
	Deadline   string `json:"deadline`
}

// get all task
func FindTasks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var tasks []models.Task
	db.Find(&tasks)
}

// create task
func CreateTasks(c *gin.Context) {
	// validate input
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// initiate data
	date := "2022-09-02"
	deadline, _ := time.Parse(date, input.Deadline)

	// create task
	task := models.Task{
		AssingedTo: input.AssingedTo,
		Task:       input.Task,
		Deadline:   deadline,
	}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&task)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// get specific task
func FindTask(c *gin.Context) {
	var task models.Task

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": task})
}

// update task
func UpdateTask(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var task models.Task
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	// Validate input
	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	date := "2022-09-05"
	deadline, _ := time.Parse(date, input.Deadline)

	var updatedInput models.Task
	updatedInput.Deadline = deadline
	updatedInput.AssingedTo = input.AssingedTo
	updatedInput.Task = input.Task

	db.Model(&task).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// delete task
func DeleteTask(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var book models.Task
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	db.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
