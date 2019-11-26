package registry

const (
	Create EventAction = iota
	Update
	Delete
)

type EventAction byte

type Registry interface {
	Register(option RegisterOption, provider ...Provider)
	UnRegister(option RegisterOption, provider ...Provider)
	GetServiceList() []Provider
	Watch() Watcher
	Unwatch(watcher Watcher)
}

type Watcher interface {
	Next() (*Event, error)
	Close()
}

type Event struct {
	AppKey    string
	Providers []Provider
}

type RegisterOption struct {
	AppKey string
}

type Provider struct {
	ProvideKey string
	Network    string
	Addr       string
	Meta       map[string]interface{}
}

type Peer2PeerDiscovery struct {
	providers []Provider
}

func (p *Peer2PeerDiscovery) Register(option RegisterOption, provider ...Provider) {
	p.providers = provider
}

func (p *Peer2PeerDiscovery) UnRegister(option RegisterOption, provider ...Provider) {
	p.providers = []Provider{}
}

func (p *Peer2PeerDiscovery) GetServiceList() []Provider {
	return p.providers
}

func (p *Peer2PeerDiscovery) Watch() Watcher {
	return nil
}

func (p *Peer2PeerDiscovery) Unwatch(watcher Watcher) {
	return
}

func (p *Peer2PeerDiscovery) WithProvider(provider Provider) *Peer2PeerDiscovery {
	p.providers = append(p.providers, provider)
	return p
}

func (p *Peer2PeerDiscovery) WithProviders(providers []Provider) *Peer2PeerDiscovery {
	for _, provider := range providers {
		p.providers = append(p.providers, provider)
	}
	return p
}

func NewPeer2PeerRegistry() *Peer2PeerDiscovery {
	r := &Peer2PeerDiscovery{}
	return r
}
