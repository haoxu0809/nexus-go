package rawrepo

import (
	"github.com/haoxu0809/nexus-go/repositorymanagement/common/model"
	"github.com/haoxu0809/nexus-go/repositorymanagement/raw/apiv1"

	"github.com/haoxu0809/pkg/log"
	"github.com/haoxu0809/pkg/rest"
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
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump()))
		return nil, err
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
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump(r)))
		return err
	}

	return nil
}

func (s *Raw) CreateRawLocalRepository(r *apiv1.RawLocalApiRequest) error {
	result := s.client.
		Verb("POST").
		SetEndpoint("service/rest/v1/repositories/raw/hosted").
		Body(r).
		Do()
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump(r)))
		return err
	}

	return nil
}

func (s *Raw) GetRawProxyRepository(repositoryName string) (*model.ProxyRepository, error) {
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/repositories/raw/proxy/" + repositoryName).
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

func (s *Raw) UpdateRawProxyRepository(repositoryName string, r *apiv1.RawProxyApiRequest) error {
	result := s.client.
		Verb("PUT").
		SetEndpoint("service/rest/v1/repositories/raw/proxy/" + repositoryName).
		Body(r).
		Do()
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump(r)))
		return err
	}

	return nil
}

func (s *Raw) CreateRawProxyRepository(r *apiv1.RawProxyApiRequest) error {
	result := s.client.
		Verb("POST").
		SetEndpoint("service/rest/v1/repositories/raw/proxy").
		Body(r).
		Do()
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump(r)))
		return err
	}

	return nil
}
