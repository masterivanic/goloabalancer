package loadbalancer

type ServerPool struct {
	Servers []*Server
}

func NewServerPool() *ServerPool {
	return &ServerPool{
		Servers: make([]*Server, 0),
	}
}

func (serverPool *ServerPool) AddServer(server *Server) error {
	serverPool.Servers = append(serverPool.Servers, server)
	return nil
}

func (serverPool *ServerPool) GetAllServers() []*Server {
	return serverPool.Servers
}
