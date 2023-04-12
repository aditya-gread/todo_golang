package routers

import (
	"main/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

var RegisterToDoRouters = func(r *gin.Engine) {

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	r.GET("/todo/", controllers.GetAllTodo)
	r.POST("/todo/", controllers.CreateTodo)
	r.GET("/todo/:id", controllers.GetTodoByID)
	r.PUT("/todo/update/:id", controllers.UpdateTodo)
	r.POST("/todo/delete/:id", controllers.DeleteTodo)
	// r.GET("/todo/search/:word", controllers.SearchTodo)
}
