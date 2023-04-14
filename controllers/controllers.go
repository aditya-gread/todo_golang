package controllers

import (
	"main/models"
	"main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// adds data
func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.BindJSON(&todo); err != nil {
		return
	}
	data := models.CreateTodo(&todo)
	c.IndentedJSON(http.StatusCreated, data)
}

// returns all data
func GetAllTodo(c *gin.Context) {
	AllTasks, err := models.GetAllToDo()
	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, AllTasks)
}

// return a data
func GetTodoByID(c *gin.Context) {
	getTodoByID, err := models.GetTodoByID(utils.GetParamInINT(c.Param("id")))
	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, getTodoByID)
}

// deletes data
func DeleteTodo(c *gin.Context) {
	models.DeleteTodo(utils.GetParamInINT(c.Param("id")))
}

// updates data
func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.BindJSON(&todo); err != nil {
		return
	}
	models.UpdateTodo(&todo, utils.GetParamInINT(c.Param("id")))
}

// search data
func SearchTodo(c *gin.Context) {
	word := c.Param(("word"))
	books := models.SearchTodo(word)

	c.IndentedJSON(http.StatusOK, books)
}
