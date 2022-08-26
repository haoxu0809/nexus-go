package search

import (
	assetsapiv1 "github.com/haoxu0809/nexus-go/assets/apiv1"
	componentsapiv1 "github.com/haoxu0809/nexus-go/components/apiv1"
	searchapiv1 "github.com/haoxu0809/nexus-go/search/apiv1"

	"github.com/haoxu0809/pkg/log"
	"github.com/haoxu0809/pkg/qencoder"
	"github.com/haoxu0809/pkg/rest"
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

func (s *Search) SearchComponents(params *searchapiv1.Components) (*componentsapiv1.ComponentList, error) {
	values, err := qencoder.Values(params)
	if err != nil {
		return nil, err
	}
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/search").
		Param(values).
		Do()
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump()))
		return nil, err
	}

	var components componentsapiv1.ComponentList
	if err := result.Into(&components); err != nil {
		return nil, err
	}

	return &components, nil
}

func (s *Search) SearchAssets(params *searchapiv1.Assets) (*assetsapiv1.AssetList, error) {
	values, err := qencoder.Values(params)
	if err != nil {
		return nil, err
	}
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/search/assets").
		Param(values).
		Do()
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump()))
		return nil, err
	}

	var assets assetsapiv1.AssetList
	if err := result.Into(&assets); err != nil {
		return nil, err
	}

	return &assets, nil
}

func (s *Search) SearchAndDownloadAsset(params *searchapiv1.DownloadAsset) error {
	values, err := qencoder.Values(params)
	if err != nil {
		return err
	}
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/search/assets/download").
		Param(values).
		Do()
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump()))
		return err
	}

	return nil
}
