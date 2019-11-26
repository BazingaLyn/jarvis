package memory

import (
	"github.com/BazingaLyn/jarvis/middleware/rpc/registry"
	"sync"
)

type Registry struct {
	mu        sync.RWMutex
	providers []registry.Provider
	watchers  map[string]*Watcher
}

func (r *Registry) Register(option registry.RegisterOption, provider ...registry.Provider) {
	r.mu.Lock()
	defer r.mu.Unlock()
	go r.sendWatcherEvent(option.AppKey, provider)

}

func (r *Registry) UnRegister(option registry.RegisterOption, provider ...registry.Provider) {
	panic("implement me")
}

func (r *Registry) GetServiceList() []registry.Provider {
	panic("implement me")
}

func (r *Registry) Watch() registry.Watcher {
	panic("implement me")
}

func (r *Registry) Unwatch(watcher registry.Watcher) {
	panic("implement me")
}

func (r *Registry) sendWatcherEvent(s string, providers []registry.Provider) {

}

type Watcher struct {
	id   string
	res  chan *registry.Event
	exit chan bool
}
