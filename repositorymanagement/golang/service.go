package golangrepo

import (
	"github.com/haoxu0809/nexus-go/repositorymanagement/common/model"
	"github.com/haoxu0809/nexus-go/repositorymanagement/golang/apiv1"

	"github.com/haoxu0809/pkg/log"
	"github.com/haoxu0809/pkg/rest"
	"go.uber.org/zap"
)

type Golang struct {
	client *rest.Client
}

func NewGolangService(client *rest.Client) *Golang {
	return &Golang{
		client: client,
	}
}

func (s *Golang) GetGolangProxyRepository(repositoryName string) (*model.ProxyRepository, error) {
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/repositories/golang/proxy/" + repositoryName).
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

func (s *Golang) UpdateGolangProxyRepository(repositoryName string, r *apiv1.GolangProxyApiRequest) error {
	result := s.client.
		Verb("PUT").SetEndpoint("service/rest/v1/repositories/golang/proxy/" + repositoryName).
		Body(r).
		Do()
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump(r)))
		return err
	}

	return nil
}

func (s *Golang) CreateGolangProxyRepository(r *apiv1.GolangProxyApiRequest) error {
	result := s.client.
		Verb("POST").
		SetEndpoint("service/rest/v1/repositories/golang/proxy").
		Body(r).
		Do()
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump(r)))
		return err
	}

	return nil
}
