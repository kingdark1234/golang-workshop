package container

import (
	"fmt"
	"workshop-service/internals/config"
	"workshop-service/internals/controller"
	"workshop-service/internals/infrastructure/database"
	"workshop-service/internals/infrastructure/http"
	"workshop-service/internals/repository"

	"go.uber.org/dig"
)

// Container ...
type Container struct {
	container  *dig.Container
	ServerBase *database.ServerBase
}

// Configure ...
func (c *Container) Configure() error {
	if err := c.container.Provide(controller.NewCalculateController); err != nil {
		return err
	}
	if err := c.container.Provide(controller.NewProblemController); err != nil {
		return err
	}
	if err := c.container.Provide(controller.NewPeopleController); err != nil {
		return err
	}
	if err := c.container.Provide(http.NewServerHTTP); err != nil {
		return err
	}
	if err := c.container.Provide(database.NewDatabaseServer); err != nil {
		return err
	}
	if err := c.container.Provide(config.NewConfiguration); err != nil {
		return err
	}
	if err := c.container.Provide(repository.NewPeopleRepository); err != nil {
		return err
	}
	return nil
}

// Start ...
func (c *Container) Start() error {
	fmt.Println("Start Container")
	if err := c.container.Invoke(func(s *http.ServerHTTP) {
		s.Start()
	}); err != nil {
		return err
	}
	return nil
}

// StartDatabase ...
func (c *Container) StartDatabase() error {
	if err := c.container.Invoke(func(db *database.ServerBase) {
		db.Initialize()
	}); err != nil {
		return err
	}
	return nil
}

// NewContainer ...
func NewContainer() (*Container, error) {
	d := dig.New()
	container := &Container{
		container: d,
	}
	if err := container.Configure(); err != nil {
		return nil, err
	}
	return container, nil
}
