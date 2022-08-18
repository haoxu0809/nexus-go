package npmrepo

import (
	"github.com/haoxu0809/nexus-go/pkg/log"
	"github.com/haoxu0809/nexus-go/pkg/rest"
	"github.com/haoxu0809/nexus-go/repositorymanagement/common/model"
	"github.com/haoxu0809/nexus-go/repositorymanagement/npm/apiv1"

	"go.uber.org/zap"
)

type Npm struct {
	client *rest.Client
}

func NewNpmService(client *rest.Client) *Npm {
	return &Npm{
		client: client,
	}
}

func (s *Npm) GetNpmLocalRepository(repositoryName string) (*model.LocalRepository, error) {
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/repositories/npm/hosted/" + repositoryName).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to get Npm local repository", zap.Object("context", log.Context(nil, result.Response())))
		return nil, result.Error()
	}

	var repository model.LocalRepository
	if err := result.Into(&repository); err != nil {
		return nil, err
	}

	return &repository, nil
}

func (s *Npm) UpdateNpmLocalRepository(repositoryName string, r *apiv1.NpmLocalApiRequest) error {
	result := s.client.
		Verb("PUT").
		SetEndpoint("service/rest/v1/repositories/npm/hosted/" + repositoryName).
		Body(r).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to update Npm local repository", zap.Object("context", log.Context(r, result.Response())))
		return result.Error()
	}

	return nil
}

func (s *Npm) CreateNpmLocalRepository(r *apiv1.NpmLocalApiRequest) error {
	result := s.client.
		Verb("POST").
		SetEndpoint("service/rest/v1/repositories/npm/hosted").
		Body(r).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to create Npm local repository", zap.Object("context", log.Context(r, result.Response())))
		return result.Error()
	}

	return nil
}

func (s *Npm) GetNpmProxyRepository(repositoryName string) (*apiv1.NpmProxyRepository, error) {
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/repositories/npm/proxy/" + repositoryName).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to get Npm proxy repository", zap.Object("context", log.Context(nil, result.Response())))
		return nil, result.Error()
	}

	var repository apiv1.NpmProxyRepository
	if err := result.Into(&repository); err != nil {
		return nil, err
	}

	return &repository, nil
}

func (s *Npm) CreateNpmProxyRepository(repositoryName string, r *apiv1.NpmProxyRepository) error {
	result := s.client.
		Verb("POST").
		SetEndpoint("service/rest/v1/repositories/npm/proxy/" + repositoryName).
		Body(r).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to create Npm proxy repository", zap.Object("context", log.Context(r, result.Response())))
		return result.Error()
	}

	return nil
}

func (s *Npm) UpdateNpmProxyRepository(r *apiv1.NpmProxyRepository) error {
	result := s.client.
		Verb("PUT").
		SetEndpoint("service/rest/v1/repositories/npm/proxy").
		Body(r).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to update Npm proxy repository", zap.Object("context", log.Context(r, result.Response())))
		return result.Error()
	}

	return nil
}
