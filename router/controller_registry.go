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

var r *controllerRegistry

func GetControllerRegistry() *controllerRegistry {
	if r == nil {
		r = &controllerRegistry{mu: sync.Mutex{}}
	}

	return r
}
