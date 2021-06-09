package models

import (
	"net/http/httputil"
	"net/url"
	"sync"
)

type Server struct {
	URL          *url.URL
	IsDead       bool
	lock         *sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}

func (server *Server) SetHealth(status bool) {
	server.lock.Lock()
	server.IsDead = status
	server.lock.Unlock()
}

func (server *Server) GetHealth() bool {
	server.lock.RLock()
	health := server.IsDead
	server.lock.RUnlock()
	return health
}

func GetNewServer(path string) *Server {
	server := new(Server)
	server.URL, _ = url.Parse(path)
	server.IsDead = false
	server.ReverseProxy = httputil.NewSingleHostReverseProxy(server.URL)
	server.lock = new(sync.RWMutex)
	return server
}
