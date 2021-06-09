package models

import "fmt"

type ServerPool struct {
	Servers        []*Server
	ServerCount    int64
	LastServerUsed int64
}

func NewServerPool() *ServerPool { return new(ServerPool) }

func (pool *ServerPool) RegisterServer(server *Server) {
	pool.Servers = append(pool.Servers, server)
	pool.ServerCount++
}

func (pool *ServerPool) GetNextAvailableServer() int64 {
	return ((pool.LastServerUsed + 1) % pool.ServerCount)
}

func (pool *ServerPool) UpdateLastServerUsed(idx int64) {
	pool.LastServerUsed = idx
	fmt.Println("last server used is updated to : ", pool.LastServerUsed)
}

func RoundRobinScheduler(pool *ServerPool) *Server {
	var counts int
	for i := pool.GetNextAvailableServer(); ; i++ {
		if counts > int(pool.ServerCount) {
			return nil
		}
		fmt.Println("i is: ", i)
		idx := (i) % pool.ServerCount
		fmt.Println("next available server is: ", idx)
		if !pool.Servers[idx].GetHealth() {
			pool.UpdateLastServerUsed(idx)
			return pool.Servers[idx]

		}
		counts++
	}
}
