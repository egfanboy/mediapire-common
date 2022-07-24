package router

import "sync"

type Controller interface {
	GetApis() []RouteBuilder
}

type ControllerRegistry struct {
	mu          sync.Mutex
	controllers []Controller
}

func (r *ControllerRegistry) Register(c Controller) {
	r.mu.Lock()

	r.controllers = append(r.controllers, c)

	r.mu.Unlock()
}

func (r *ControllerRegistry) GetControllers() []Controller {
	return r.controllers
}

func NewControllerRegistry() *ControllerRegistry {
	return &ControllerRegistry{mu: sync.Mutex{}}
}
