package router

import "sync"

type Controller interface {
	GetApis() []RouteBuilder
}

type controllerRegistry struct {
	mu          sync.Mutex
	controllers []Controller
}

func (r *controllerRegistry) Register(c Controller) {
	r.mu.Lock()

	r.controllers = append(r.controllers, c)

	r.mu.Unlock()
}

func (r *controllerRegistry) GetControllers() []Controller {
	return r.controllers
}

func NewControllerRegistry() *controllerRegistry {
	return &controllerRegistry{mu: sync.Mutex{}}
}
