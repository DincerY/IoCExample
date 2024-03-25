package src

type Container struct {
	services map[string]string
}

func (container *Container) AddService(abstract string, concrete string) {
	if container == nil {
		panic("Services is not exist")
	}
	container.services[abstract] = concrete

}

func NewContainer() *Container {
	return &Container{
		services: make(map[string]string),
	}
}

func (container *Container) GetServices() map[string]string {
	return container.services
}

func (container *Container) GetService(abstract string) string {
	return container.services[abstract]
}
