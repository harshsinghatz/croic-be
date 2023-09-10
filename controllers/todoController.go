package controllers

import (
	"croic/initializers"
	"croic/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddTodo(c *gin.Context) {
	var user, exists = c.Get("user")

	userModal, ok := user.(models.User)

	if !exists && !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong",
		})

		return
	}

	var body struct {
		Title       string
		Description string
		DueDate     string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong",
		})

		return
	}

	newTodo := models.Todo{
		Title:       body.Title,
		Description: body.Description,
		Status:      "planned",
		DueDate:     body.DueDate,
		UserID:      userModal.ID,
	}

	initializers.DB.Create(&newTodo)

	c.JSON(http.StatusOK, gin.H{
		"message": "Todo successfully created!",
	})
}

func GetAllTodos(c *gin.Context) {
	var user, exists = c.Get("user")

	userModal, ok := user.(models.User)

	if !exists && !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong",
		})

		return
	}

	// var todos []models.Todo

	// result := initializers.DB.Where("user_id = ?", userModal.ID).Find(todos)

	todos := make([]models.Todo, 0) // Initialize the slice

	// Retrieve todos associated with the given user ID
	result := initializers.DB.Where("user_id = ?", userModal.ID).Find(&todos)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Todos retrieved succesfully",
		"data":    todos,
	})
}

func UpdateTodoById(c *gin.Context) {

	// var user, exists = c.Get("user")

	// userModal, ok := user.(models.User)

	// if !exists && !ok {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "Something went wrong",
	// 	})

	// 	return
	// }

	var todoId = c.Param("id")

	var todo models.Todo

	initializers.DB.Where("id = ?", todoId).Find(&todo)

	c.JSON(http.StatusBadRequest, gin.H{
		"message": "Something went wrong",
	})

	return

}
