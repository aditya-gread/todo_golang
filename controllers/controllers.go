package controllers

import (
	"main/models"
	"main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.BindJSON(&todo); err != nil {
		return
	}
	models.CreateTodo(&todo)
	c.IndentedJSON(http.StatusCreated, todo)
}

func GetAllTodo(c *gin.Context) {
	AllTasks, err := models.GetAllToDo()
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, AllTasks)
}

func GetTodoByID(c *gin.Context) {
	getTodoByID, err := models.GetTodoByID(utils.GetParamInINT(c.Param("id")))
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, getTodoByID)
}

func DeleteTodo(c *gin.Context) {
	models.DeleteTodo(utils.GetParamInINT(c.Param("id")))
}

func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.BindJSON(&todo); err != nil {
		panic(err)
	}
	models.UpdateTodo(&todo, utils.GetParamInINT(c.Param("id")))
}
