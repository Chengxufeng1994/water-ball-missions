package registry

import (
	"sync"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/prescriber_system/internal/domain"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/prescriber_system/internal/domain/prescribehandler"
)

type Registry struct {
	handlers map[string]domain.PrescribeHandler
	mu       sync.RWMutex
}

var _ domain.PrescribeHandlerRegistry = (*Registry)(nil)

func NewPrescribeHandlerRegistry() *Registry {
	reg := &Registry{
		handlers: make(map[string]domain.PrescribeHandler),
	}

	reg.Register("COVID-19", prescribehandler.NewCovid19PrescribeHandler())
	reg.Register("Attractive", prescribehandler.NewAttractivePrescribeHandler())
	reg.Register("SleepApneaSyndrome", prescribehandler.NewSleepApneaSyndromePrescribeHandler())

	return reg
}

// Find implements domain.PrescribeHandlerRegistry.
func (r *Registry) Find(name string) domain.PrescribeHandler {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.handlers[name]
}

// Register implements domain.PrescribeHandlerRegistry.
func (r *Registry) Register(name string, prescribeHandler domain.PrescribeHandler) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.handlers[name] = prescribeHandler
}
