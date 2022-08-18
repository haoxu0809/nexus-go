package search

import (
	assetsapiv1 "github.com/haoxu0809/nexus-go/assets/apiv1"
	componentsapiv1 "github.com/haoxu0809/nexus-go/components/apiv1"
	"github.com/haoxu0809/nexus-go/pkg/log"
	"github.com/haoxu0809/nexus-go/pkg/rest"
	"github.com/haoxu0809/nexus-go/pkg/xurl"
	searchapiv1 "github.com/haoxu0809/nexus-go/search/apiv1"

	"go.uber.org/zap"
)

type Search struct {
	client *rest.Client
}

func NewSearchService(client *rest.Client) *Search {
	return &Search{
		client: client,
	}
}

func (s *Search) SearchComponents(params *searchapiv1.SearchComponentsParameters) (*componentsapiv1.ComponentList, error) {
	values, err := xurl.Values(params)
	if err != nil {
		return nil, err
	}
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/search").
		Param(values).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to search", zap.Object("context", log.Context(nil, result.Response())))
		return nil, result.Error()
	}

	var components componentsapiv1.ComponentList
	if err := result.Into(&components); err != nil {
		return nil, err
	}

	return &components, nil
}

func (s *Search) SearchAssets(params *searchapiv1.SearchAssetsParameters) (*assetsapiv1.AssetList, error) {
	values, err := xurl.Values(params)
	if err != nil {
		return nil, err
	}
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/search/assets").
		Param(values).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to search", zap.Object("context", log.Context(nil, result.Response())))
		return nil, result.Error()
	}

	var assets assetsapiv1.AssetList
	if err := result.Into(&assets); err != nil {
		return nil, err
	}

	return &assets, nil
}

func (s *Search) SearchAndDownloadAsset(params *searchapiv1.SearchAndDownloadAssetParameters) error {
	values, err := xurl.Values(params)
	if err != nil {
		return err
	}
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/search/assets/download").
		Param(values).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to search", zap.Object("context", log.Context(nil, result.Response())))
		return result.Error()
	}

	return nil
}
