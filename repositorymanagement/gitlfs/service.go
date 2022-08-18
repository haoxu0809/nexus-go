package gitlfsrepo

import (
	"github.com/haoxu0809/nexus-go/repositorymanagement/common/model"
	"github.com/haoxu0809/nexus-go/repositorymanagement/gitlfs/apiv1"

	"github.com/haoxu0809/pkg/log"
	"github.com/haoxu0809/pkg/rest"
	"go.uber.org/zap"
)

type GitLfs struct {
	client *rest.Client
}

func NewGitLfsService(client *rest.Client) *GitLfs {
	return &GitLfs{
		client: client,
	}
}

func (s *GitLfs) GetGitLfsLocalRepository(repositoryName string) (*model.LocalRepository, error) {
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/repositories/gitlfs/hosted/" + repositoryName).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to get Git LFS local repository", zap.Object("context", log.Context(nil, result.Response())))
		return nil, result.Error()
	}

	var repository model.LocalRepository
	if err := result.Into(&repository); err != nil {
		return nil, err
	}

	return &repository, nil
}

func (s *GitLfs) UpdateGitLfsLocalRepository(repositoryName string, r *apiv1.GitLfsLocalApiRequest) error {
	result := s.client.
		Verb("PUT").
		SetEndpoint("service/rest/v1/repositories/gitlfs/hosted/" + repositoryName).
		Body(r).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to update Git LFS local repository", zap.Object("context", log.Context(r, result.Response())))
		return result.Error()
	}

	return nil
}

func (s *GitLfs) CreateGitLfsLocalRepository(r *apiv1.GitLfsLocalApiRequest) error {
	result := s.client.
		Verb("POST").
		SetEndpoint("service/rest/v1/repositories/gitlfs/hosted").
		Body(r).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to create Git LFS local repository", zap.Object("context", log.Context(r, result.Response())))
		return result.Error()
	}

	return nil
}
