package pool

import "load-balancer/models"

var serverPool *models.ServerPool

func InitServerPool() {
	serverPool = models.NewServerPool()
	server1 := models.GetNewServer("http://localhost:7000")
	server2 := models.GetNewServer("http://localhost:7000")
	serverPool.RegisterServer(server1)
	serverPool.RegisterServer(server2)
}

func GetServerPool() *models.ServerPool { return serverPool }
