package golangrepo

import (
	"github.com/haoxu0809/nexus-go/pkg/log"
	"github.com/haoxu0809/nexus-go/pkg/rest"
	"github.com/haoxu0809/nexus-go/repositorymanagement/common/model"
	"github.com/haoxu0809/nexus-go/repositorymanagement/golang/apiv1"

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
	if result.Error() != nil {
		log.L().Error("failed to get Golang proxy repository", zap.Object("context", log.Context(nil, result.Response())))
		return nil, result.Error()
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
	if result.Error() != nil {
		log.L().Error("failed to update Golang proxy repository", zap.Object("context", log.Context(r, result.Response())))
		return result.Error()
	}

	return nil
}

func (s *Golang) CreateGolangProxyRepository(r *apiv1.GolangProxyApiRequest) error {
	result := s.client.
		Verb("POST").
		SetEndpoint("service/rest/v1/repositories/golang/proxy").
		Body(r).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to create Golang proxy repository", zap.Object("context", log.Context(r, result.Response())))
		return result.Error()
	}

	return nil
}