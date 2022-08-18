package rawrepo

import (
	"github.com/haoxu0809/nexus-go/pkg/log"
	"github.com/haoxu0809/nexus-go/pkg/rest"
	"github.com/haoxu0809/nexus-go/repositorymanagement/common/model"
	"github.com/haoxu0809/nexus-go/repositorymanagement/raw/apiv1"

	"go.uber.org/zap"
)

type Raw struct {
	client *rest.Client
}

func NewRawService(client *rest.Client) *Raw {
	return &Raw{
		client: client,
	}
}

func (s *Raw) GetRawLocalRepository(repositoryName string) (*apiv1.RawLocalRepository, error) {
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/repositories/raw/hosted/" + repositoryName).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to get raw local repository", zap.Object("context", log.Context(nil, result.Response())))
		return nil, result.Error()
	}

	var repository apiv1.RawLocalRepository
	if err := result.Into(&repository); err != nil {
		return nil, err
	}

	return &repository, nil
}

func (s *Raw) UpdateRawLocalRepository(repositoryName string, r *apiv1.RawLocalApiRequest) error {
	result := s.client.
		Verb("PUT").
		SetEndpoint("service/rest/v1/repositories/raw/hosted/" + repositoryName).
		Body(r).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to update raw local repository", zap.Object("context", log.Context(r, result.Response())))
		return result.Error()
	}

	return nil
}

func (s *Raw) CreateRawLocalRepository(r *apiv1.RawLocalApiRequest) error {
	result := s.client.
		Verb("POST").
		SetEndpoint("service/rest/v1/repositories/raw/hosted").
		Body(r).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to create raw local repository", zap.Object("context", log.Context(r, result.Response())))
		return result.Error()
	}

	return nil
}

func (s *Raw) GetRawProxyRepository(repositoryName string) (*model.ProxyRepository, error) {
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/repositories/raw/proxy/" + repositoryName).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to get raw proxy repository", zap.Object("context", log.Context(nil, result.Response())))
		return nil, result.Error()
	}

	var repository model.ProxyRepository
	if err := result.Into(&repository); err != nil {
		return nil, err
	}

	return &repository, nil
}

func (s *Raw) UpdateRawProxyRepository(repositoryName string, r *apiv1.RawProxyApiRequest) error {
	result := s.client.
		Verb("PUT").
		SetEndpoint("service/rest/v1/repositories/raw/proxy/" + repositoryName).
		Body(r).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to update raw proxy repository", zap.Object("context", log.Context(r, result.Response())))
		return result.Error()
	}

	return nil
}

func (s *Raw) CreateRawProxyRepository(r *apiv1.RawProxyApiRequest) error {
	result := s.client.
		Verb("POST").
		SetEndpoint("service/rest/v1/repositories/raw/proxy").
		Body(r).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to create raw proxy repository", zap.Object("context", log.Context(r, result.Response())))
		return result.Error()
	}

	return nil
}
