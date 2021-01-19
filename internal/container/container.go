package container

import (
	"5g-v2x-data-management-service/internal/config"
	"5g-v2x-data-management-service/internal/infrastructures/database"
	"5g-v2x-data-management-service/internal/infrastructures/kafka"
	"5g-v2x-data-management-service/internal/repositories"
	"5g-v2x-data-management-service/internal/services"

	"go.uber.org/dig"
)

// Container ...
type Container struct {
	container *dig.Container
	Error     error
}

// NewContainer ...
func NewContainer() *Container {
	c := new(Container)
	c.Configure()
	return c
}

// Configure ...
func (cn *Container) Configure() {
	cn.container = dig.New()

	if err := cn.container.Provide(config.NewConfig); err != nil {
		cn.Error = err
	}

	if err := cn.container.Provide(kafka.NewConsumer); err != nil {
		cn.Error = err
	}

	if err := cn.container.Provide(database.NewMongoDatabase); err != nil {
		cn.Error = err
	}

	if err := cn.container.Provide(services.NewAccidentService); err != nil {
		cn.Error = err
	}
	if err := cn.container.Provide(services.NewDrowsinessService); err != nil {
		cn.Error = err
	}

	if err := cn.container.Provide(repositories.NewCRUDRepository); err != nil {
		cn.Error = err
	}

}

// Run ...
func (cn *Container) Run() *Container {
	if err := cn.container.Invoke(func(c *kafka.Consumer) {
		if err := c.Start(); err != nil {
			panic(err)
		}
	}); err != nil {
		panic(err)
	}
	return cn
}
