package helmrepo

import (
	"github.com/haoxu0809/nexus-go/pkg/log"
	"github.com/haoxu0809/nexus-go/pkg/rest"
	"github.com/haoxu0809/nexus-go/repositorymanagement/common/model"
	"github.com/haoxu0809/nexus-go/repositorymanagement/helm/apiv1"

	"go.uber.org/zap"
)

type Helm struct {
	client *rest.Client
}

func NewHelmService(client *rest.Client) *Helm {
	return &Helm{
		client: client,
	}
}

func (s *Helm) GetHelmLocalRepository(repositoryName string) (*model.LocalRepository, error) {
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/repositories/helm/hosted/" + repositoryName).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to get Helm local repository", zap.Object("context", log.Context(nil, result.Response())))
		return nil, result.Error()
	}

	var repository model.LocalRepository
	if err := result.Into(&repository); err != nil {
		return nil, err
	}

	return &repository, nil
}

func (s *Helm) UpdateHelmLocalRepository(repositoryName string, r *apiv1.HelmLocalApiRequest) error {
	result := s.client.
		Verb("PUT").
		SetEndpoint("service/rest/v1/repositories/helm/hosted/" + repositoryName).
		Body(r).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to update Helm local repository", zap.Object("context", log.Context(r, result.Response())))
		return result.Error()
	}

	return nil
}

func (s *Helm) CreateHelmLocalRepository(r *apiv1.HelmLocalApiRequest) error {
	result := s.client.
		Verb("POST").
		SetEndpoint("service/rest/v1/repositories/helm/hosted").
		Body(r).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to create Helm local repository", zap.Object("context", log.Context(r, result.Response())))
		return result.Error()
	}

	return nil
}

func (s *Helm) GetHelmProxyRepository(repositoryName string) (*model.ProxyRepository, error) {
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/repositories/helm/proxy/" + repositoryName).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to get Helm proxy repository", zap.Object("context", log.Context(nil, result.Response())))
		return nil, result.Error()
	}

	var repository model.ProxyRepository
	if err := result.Into(&repository); err != nil {
		return nil, err
	}

	return &repository, nil
}

func (s *Helm) UpdateHelmProxyRepository(repositoryName string, r *apiv1.HelmProxyApiRequest) error {
	result := s.client.
		Verb("PUT").
		SetEndpoint("service/rest/v1/repositories/helm/proxy/" + repositoryName).
		Body(r).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to update Helm proxy repository", zap.Object("context", log.Context(r, result.Response())))
		return result.Error()
	}

	return nil
}

func (s *Helm) CreateHelmProxyRepository(r *apiv1.HelmProxyApiRequest) error {
	result := s.client.
		Verb("POST").
		SetEndpoint("service/rest/v1/repositories/helm/proxy").
		Body(r).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to create Helm proxy repository", zap.Object("context", log.Context(r, result.Response())))
		return result.Error()
	}

	return nil
}
