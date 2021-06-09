package main

import (
	"load-balancer/pool"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	InitRoutes(router)
	pool.InitServerPool()
	router.Run(":8080")
}
