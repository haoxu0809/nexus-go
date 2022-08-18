package repositories

import (
	"github.com/haoxu0809/nexus-go/repositorymanagement/repositories/apiv1"

	"github.com/haoxu0809/pkg/log"
	"github.com/haoxu0809/pkg/rest"
	"go.uber.org/zap"
)

type Repository struct {
	client *rest.Client
}

func NewRepositoryService(client *rest.Client) *Repository {
	return &Repository{
		client: client,
	}
}

func (s *Repository) InvalidateCache(repositoryName string) error {
	result := s.client.
		Verb("POST").
		SetEndpoint("service/rest/v1/repositories/" + repositoryName + "/invalidate-cache").
		Do()
	if result.Error() != nil {
		log.L().Error("failed to invalidate cache", zap.Object("context", log.Context(nil, result.Response())))
		return result.Error()
	}

	return nil
}

func (s *Repository) RebuildIndex(repositoryName string) error {
	result := s.client.
		Verb("POST").
		SetEndpoint("service/rest/v1/repositories/" + repositoryName + "/rebuild-index").
		Do()
	if result.Error() != nil {
		log.L().Error("failed to rebuild index", zap.Object("context", log.Context(nil, result.Response())))
		return result.Error()
	}

	return nil
}

func (s *Repository) GetRepository(repositoryName string) (*apiv1.Repository, error) {
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/repositories/" + repositoryName).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to get repository", zap.Object("context", log.Context(nil, result.Response())))
		return nil, result.Error()
	}

	var repository apiv1.Repository
	if err := result.Into(&repository); err != nil {
		return nil, err
	}

	return &repository, nil
}

func (s *Repository) DeleteRepository(repositoryName string) error {
	result := s.client.
		Verb("DELETE").
		SetEndpoint("service/rest/v1/repositories/" + repositoryName).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to delete repository", zap.Object("context", log.Context(nil, result.Response())))
		return result.Error()
	}

	return nil
}

func (s *Repository) ListRepository() ([]*apiv1.Repository, error) {
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/repositories").
		Do()
	if result.Error() != nil {
		log.L().Error("failed to list repositories", zap.Object("context", log.Context(nil, result.Response())))
		return nil, result.Error()
	}

	var repositories []*apiv1.Repository
	if err := result.Into(&repositories); err != nil {
		return nil, err
	}

	return repositories, nil
}

func (s *Repository) ListRepositorySettings() ([]*apiv1.RepositorySettings, error) {
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/repositorySettings").
		Do()
	if result.Error() != nil {
		log.L().Error("failed to list repository settings", zap.Object("context", log.Context(nil, result.Response())))
		return nil, result.Error()
	}

	var settings []*apiv1.RepositorySettings
	if err := result.Into(&settings); err != nil {
		return nil, err
	}

	return settings, nil
}
