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
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump()))
		return err
	}

	return nil
}

func (s *Repository) RebuildIndex(repositoryName string) error {
	result := s.client.
		Verb("POST").
		SetEndpoint("service/rest/v1/repositories/" + repositoryName + "/rebuild-index").
		Do()
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump()))
		return err
	}

	return nil
}

func (s *Repository) GetRepository(format, rtype, repositoryName string) (*apiv1.Repository, error) {
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/repositories/" + format + "/" + rtype + "/" + repositoryName).
		Do()
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump()))
		return nil, err
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
	if result.StatusCode() == 404 {
		return nil
	}
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump()))
		return err
	}

	return nil
}

func (s *Repository) ListRepository() ([]*apiv1.Repository, error) {
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/repositories").
		Do()
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump()))
		return nil, err
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
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump()))
		return nil, err
	}

	var settings []*apiv1.RepositorySettings
	if err := result.Into(&settings); err != nil {
		return nil, err
	}

	return settings, nil
}
