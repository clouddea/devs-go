package modeling

type Coupled interface {
	Entity
	AddComponent(entity Entity)
	AddCoupling(from Entity, fromPort string, to Entity, toPort string)
	GetComponents() []Entity
	GetComponentMap() map[string]Entity
	GetCoupling(from Entity, port string) []Pair
}

type Pair struct {
	Component Entity
	Port      string
}

type AbstractCoupled struct {
	components map[string]Entity
	couplings  map[string]map[string][]Pair
	name       string
}

func (receiver AbstractCoupled) Name() string {
	return receiver.name
}

func (receiver *AbstractCoupled) SetName(name string) {
	receiver.name = name
}

func (receiver *AbstractCoupled) AddComponent(entity Entity) {
	if receiver.components == nil {
		receiver.components = make(map[string]Entity)
	}
	receiver.components[entity.Name()] = entity
}

func (receiver *AbstractCoupled) AddCoupling(from Entity, fromPort string, to Entity, toPort string) {
	// 当组件不是子组件也不是当前组件时，忽略
	if _, ok := receiver.components[from.Name()]; !ok && from.Name() != receiver.Name() {
		return
	}
	if _, ok := receiver.components[to.Name()]; !ok && to.Name() != receiver.Name() {
		return
	}
	// 添加耦合关系
	if receiver.couplings == nil {
		receiver.couplings = make(map[string]map[string][]Pair)
	}
	portMap := receiver.couplings[from.Name()]
	if portMap == nil {
		portMap = make(map[string][]Pair)
		receiver.couplings[from.Name()] = portMap
	}

	val := Pair{to, toPort}
	portMap[fromPort] = append(portMap[fromPort], val)
}

func (receiver *AbstractCoupled) GetComponents() []Entity {
	var components []Entity
	for _, v := range receiver.components {
		components = append(components, v)
	}
	return components
}

func (receiver *AbstractCoupled) GetComponentMap() map[string]Entity {
	if receiver.components == nil {
		receiver.components = make(map[string]Entity)
	}
	return receiver.components
}

func (receiver *AbstractCoupled) GetCoupling(from Entity, fromPort string) (results []Pair) {
	results = make([]Pair, 0)
	var portMap = receiver.couplings[from.Name()]
	if portMap == nil {
		return
	}
	results = portMap[fromPort]
	return
}
