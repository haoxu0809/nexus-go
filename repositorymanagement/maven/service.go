package mavenrepo

import (
	"github.com/haoxu0809/nexus-go/repositorymanagement/maven/apiv1"

	"github.com/haoxu0809/pkg/log"
	"github.com/haoxu0809/pkg/rest"
	"go.uber.org/zap"
)

type Maven struct {
	client *rest.Client
}

func NewMavenService(client *rest.Client) *Maven {
	return &Maven{
		client: client,
	}
}

func (s *Maven) GetMavenLocalRepository(repositoryName string) (*apiv1.MavenLocalRepository, error) {
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/repositories/maven/hosted/" + repositoryName).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to get maven local repository", zap.Object("context", log.Context(nil, result.Response())))
		return nil, result.Error()
	}

	var repository apiv1.MavenLocalRepository
	if err := result.Into(&repository); err != nil {
		return nil, err
	}

	return &repository, nil
}

func (s *Maven) UpdateMavenLocalRepository(repositoryName string, r *apiv1.MavenLocalApiRequest) error {
	result := s.client.
		Verb("PUT").
		SetEndpoint("service/rest/v1/repositories/maven/hosted/" + repositoryName).
		Body(r).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to update maven local repository", zap.Object("context", log.Context(r, result.Response())))
		return result.Error()
	}

	return nil
}

func (s *Maven) CreateMavenLocalRepository(r *apiv1.MavenLocalApiRequest) error {
	result := s.client.
		Verb("POST").
		SetEndpoint("service/rest/v1/repositories/maven/hosted").
		Body(r).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to create maven local repository", zap.Object("context", log.Context(r, result.Response())))
		return result.Error()
	}

	return nil
}

func (s *Maven) GetMavenProxyRepository(repositoryName string) (*apiv1.MavenProxyRepository, error) {
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/repositories/maven/proxy/" + repositoryName).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to get maven proxy repository", zap.Object("context", log.Context(nil, result.Response())))
		return nil, result.Error()
	}

	var repository apiv1.MavenProxyRepository
	if err := result.Into(&repository); err != nil {
		return nil, err
	}

	return &repository, nil
}

func (s *Maven) UpdateMavenProxyRepository(repositoryName string, r *apiv1.MavenProxyApiRequest) error {
	result := s.client.
		Verb("PUT").
		SetEndpoint("service/rest/v1/repositories/maven/proxy/" + repositoryName).
		Body(r).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to update maven proxy repository", zap.Object("context", log.Context(r, result.Response())))
		return result.Error()
	}

	return nil
}

func (s *Maven) CreateMavenProxyRepository(r *apiv1.MavenProxyApiRequest) error {
	result := s.client.
		Verb("POST").
		SetEndpoint("service/rest/v1/repositories/maven/proxy").
		Body(r).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to create maven proxy repository", zap.Object("context", log.Context(r, result.Response())))
		return result.Error()
	}

	return nil
}
