package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	InitRoutes(router)
	router.Run(":8080")
}
