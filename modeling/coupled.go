package modeling

type Coupled interface {
	AddComponent(atomic Atomic)
	AddCoupling(from Atomic, fromPort string, to Atomic, toPort string)
	GetComponents() []Atomic
	GetCoupling(from Atomic) []Atomic
}

type pair struct {
	component Atomic
	port      string
}

type AbstractCoupled struct {
	components []Atomic
	couplings  map[Atomic]map[string][]pair
	name       string
}

func (receiver *AbstractCoupled) AddComponent(atomic Atomic) {
	receiver.components = append(receiver.components, atomic)
}

func (receiver *AbstractCoupled) AddCoupling(from Atomic, fromPort string, to Atomic, toPort string) {
	if receiver.couplings == nil {
		receiver.couplings = make(map[Atomic]map[string][]pair)
	}
	portMap := receiver.couplings[from]
	if portMap == nil {
		portMap = make(map[string][]pair)
		receiver.couplings[from] = portMap
	}

	val := pair{to, toPort}
	portMap[fromPort] = append(portMap[fromPort], val)
}

func (receiver *AbstractCoupled) GetComponents() []Atomic {
	return receiver.components
}

func (receiver *AbstractCoupled) GetCoupling(from Atomic) []Atomic {
	var results []Atomic = nil
	if receiver.couplings[from] != nil {
		var portMap = receiver.couplings[from]
		for _, pairs := range portMap {
			for _, p := range pairs {
				results = append(results, p.component)
			}
		}
	}
	return results
}
