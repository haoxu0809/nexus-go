package pypirepo

import (
	"github.com/haoxu0809/nexus-go/repositorymanagement/common/model"
	"github.com/haoxu0809/nexus-go/repositorymanagement/pypi/apiv1"

	"github.com/haoxu0809/pkg/log"
	"github.com/haoxu0809/pkg/rest"
	"go.uber.org/zap"
)

type Pypi struct {
	client *rest.Client
}

func NewPypiService(client *rest.Client) *Pypi {
	return &Pypi{
		client: client,
	}
}

func (s *Pypi) GetPypiLocalRepository(repositoryName string) (*model.LocalRepository, error) {
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/repositories/pypi/hosted/" + repositoryName).
		Do()
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump()))
		return nil, err
	}

	var repository model.LocalRepository
	if err := result.Into(&repository); err != nil {
		return nil, err
	}

	return &repository, nil
}

func (s *Pypi) UpdatePypiLocalRepository(repositoryName string, r *apiv1.PypiLocalApiRequest) error {
	result := s.client.
		Verb("PUT").
		SetEndpoint("service/rest/v1/repositories/pypi/hosted/" + repositoryName).
		Body(r).
		Do()
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump(r)))
		return err
	}

	return nil
}

func (s *Pypi) CreatePypiLocalRepository(r *apiv1.PypiLocalApiRequest) error {
	result := s.client.
		Verb("POST").
		SetEndpoint("service/rest/v1/repositories/pypi/hosted").
		Body(r).
		Do()
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump(r)))
		return err
	}

	return nil
}

func (s *Pypi) GetPypiProxyRepository(repositoryName string) (*model.ProxyRepository, error) {
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/repositories/pypi/proxy/" + repositoryName).
		Do()
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump()))
		return nil, err
	}

	var repository model.ProxyRepository
	if err := result.Into(&repository); err != nil {
		return nil, err
	}

	return &repository, nil
}

func (s *Pypi) UpdatePypiProxyRepository(repositoryName string, r *apiv1.PypiProxyApiRequest) error {
	result := s.client.
		Verb("PUT").
		SetEndpoint("service/rest/v1/repositories/pypi/proxy/" + repositoryName).
		Body(r).
		Do()
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump(r)))
		return err
	}

	return nil
}

func (s *Pypi) CreatePypiProxyRepository(r *apiv1.PypiProxyApiRequest) error {
	result := s.client.
		Verb("POST").
		SetEndpoint("service/rest/v1/repositories/pypi/proxy").
		Body(r).
		Do()
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump(r)))
		return err
	}

	return nil
}
