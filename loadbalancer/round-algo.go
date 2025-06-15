package loadbalancer

import (
	"errors"
	"sync"
)

var (
	once = sync.Once{}
)

type RoundRobinBalancer struct {
	pool *ServerPool
	mu   sync.Mutex
	idx  int
}

func NewRoundRobinBalancer(pool *ServerPool) *RoundRobinBalancer {
	return &RoundRobinBalancer{
		pool: pool,
		idx:  -1,
	}
}

func (rb *RoundRobinBalancer) GetNextServer() (*Server, error) {
	servers := rb.pool.GetAllServers()

	if len(servers) == 0 {
		return nil, errors.New("no servers found")
	}

	rb.mu.Lock()
	defer rb.mu.Unlock()

	rb.idx = (rb.idx + 1) % len(servers)

	selectedServer := servers[rb.idx]

	return selectedServer, nil
}
