package dockerrepo

import (
	"github.com/haoxu0809/nexus-go/repositorymanagement/docker/apiv1"

	"github.com/haoxu0809/pkg/log"
	"github.com/haoxu0809/pkg/rest"
	"go.uber.org/zap"
)

type Docker struct {
	client *rest.Client
}

func NewDockerService(client *rest.Client) *Docker {
	return &Docker{
		client: client,
	}
}

func (s *Docker) GetDockerLocalRepository(repositoryName string) (*apiv1.DockerLocalRepository, error) {
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/repositories/docker/hosted/" + repositoryName).
		Do()
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump()))
		return nil, err
	}

	var repository apiv1.DockerLocalRepository
	if err := result.Into(&repository); err != nil {
		return nil, err
	}

	return &repository, nil
}

func (s *Docker) UpdateDockerLocalRepository(repositoryName string, r *apiv1.DockerLocalApiRequest) error {
	result := s.client.
		Verb("PUT").
		SetEndpoint("service/rest/v1/repositories/docker/hosted/" + repositoryName).
		Body(r).
		Do()
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump(r)))
		return err
	}

	return nil
}

func (s *Docker) CreateDockerLocalRepository(r *apiv1.DockerLocalApiRequest) error {
	result := s.client.
		Verb("POST").
		SetEndpoint("service/rest/v1/repositories/docker/hosted").
		Body(r).
		Do()
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump(r)))
		return err
	}

	return nil
}

func (s *Docker) GetDockerProxyRepository(repositoryName string) (*apiv1.DockerProxyRepository, error) {
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/repositories/docker/proxy/" + repositoryName).
		Do()
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump()))
		return nil, err
	}

	var repository apiv1.DockerProxyRepository
	if err := result.Into(&repository); err != nil {
		return nil, err
	}

	return &repository, nil
}

func (s *Docker) UpdateDockerProxyRepository(repositoryName string, r *apiv1.DockerProxyApiRequest) error {
	result := s.client.
		Verb("PUT").
		SetEndpoint("service/rest/v1/repositories/docker/proxy/" + repositoryName).
		Body(r).
		Do()
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump(r)))
		return err
	}

	return nil
}

func (s *Docker) CreateDockerProxyRepository(r *apiv1.DockerProxyApiRequest) error {
	result := s.client.
		Verb("POST").
		SetEndpoint("service/rest/v1/repositories/docker/proxy").
		Body(r).
		Do()
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump(r)))
		return err
	}

	return nil
}
