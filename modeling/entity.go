package modeling

type Entity interface {
	Name() string
}

type EntityRemote struct {
	name     string
	endpoint string
}

func NewEntityRemote(name string, endpoint string) *EntityRemote {
	return &EntityRemote{
		name:     name,
		endpoint: endpoint,
	}
}

func (receiver EntityRemote) Name() string {
	return receiver.name
}

func (receiver EntityRemote) EndPoint() string {
	return receiver.endpoint
}
