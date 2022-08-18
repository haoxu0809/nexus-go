package cocoapodsrepo

import (
	"github.com/haoxu0809/nexus-go/repositorymanagement/cocoapods/apiv1"
	"github.com/haoxu0809/nexus-go/repositorymanagement/common/model"

	"github.com/haoxu0809/pkg/log"
	"github.com/haoxu0809/pkg/rest"
	"go.uber.org/zap"
)

type CocoaPods struct {
	client *rest.Client
}

func NewCocoaPodsService(client *rest.Client) *CocoaPods {
	return &CocoaPods{
		client: client,
	}
}

func (s *CocoaPods) GetCocoaPodsProxyRepository(repositoryName string) (*model.ProxyRepository, error) {
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/repositories/cocoapods/proxy/" + repositoryName).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to get CocoaPods proxy repository", zap.Object("context", log.Context(nil, result.Response())))
		return nil, result.Error()
	}

	var repository model.ProxyRepository
	if err := result.Into(&repository); err != nil {
		return nil, err
	}

	return &repository, nil
}

func (s *CocoaPods) UpdateCocoaPodsProxyRepository(repositoryName string, r *apiv1.CocoapodsProxyApiRequest) error {
	result := s.client.
		Verb("PUT").
		SetEndpoint("service/rest/v1/repositories/cocoapods/proxy/" + repositoryName).
		Body(r).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to update CocoaPods proxy repository", zap.Object("context", log.Context(nil, result.Response())))
		return result.Error()
	}

	return nil
}

func (s *CocoaPods) CreateCocoaPodsProxyRepository(r *apiv1.CocoapodsProxyApiRequest) error {
	result := s.client.
		Verb("POST").
		SetEndpoint("service/rest/v1/repositories/cocoapods/proxy/").
		Body(r).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to create CocoaPods proxy repository", zap.Object("context", log.Context(r, result.Response())))
		return result.Error()
	}

	return nil
}
