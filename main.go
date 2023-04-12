package main

import (
	"main/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	routers.RegisterToDoRouters(r)
	r.Run("localhost:12345")
}
