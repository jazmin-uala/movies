package testutils

import (
	"fmt"
	"github.com/ory/dockertest/v3"
	"net/http"
)

type DockerContainer struct {
	Pool       *dockertest.Pool
	Resource   *dockertest.Resource
	MappedPort string
	Endpoint   string
}

func SetupContainer(c ContainerConfiguration) (*DockerContainer, error) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, err
	}

	resource, err := pool.Run(c.Image, c.Tag, c.Env)
	if err != nil {
		return nil, err
	}

	port := resource.GetPort(c.Port)
	if err := pool.Retry(func() error { //wait for container to startup
		_, err := http.Get(fmt.Sprintf("http://%s:%s", "localhost", port))
		if err != nil {
			return err
		} else {
			fmt.Println(fmt.Sprintf("container %s started", resource.Container.ID))
			return nil
		}
	}); err != nil {
		return nil, err
	}

	return &DockerContainer{
		Pool:       pool,
		Resource:   resource,
		MappedPort: port,
		Endpoint:   fmt.Sprintf("http://%s:%s", "localhost", port),
	}, nil
}

type ContainerConfiguration struct {
	Image string
	Tag   string
	Port  string
	Env   []string
}

func CreateDynamoContainer() (*DockerContainer, error) {
	return SetupContainer(ContainerConfiguration{
		Image: "amazon/dynamodb-local",
		Tag:   "latest",
		Port:  "8000/tcp",
		Env:   nil,
	})
}







