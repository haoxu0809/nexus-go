package assets

import (
	"github.com/haoxu0809/nexus-go/assets/apiv1"

	"github.com/haoxu0809/pkg/log"
	"github.com/haoxu0809/pkg/qencoder"
	"github.com/haoxu0809/pkg/rest"
	"go.uber.org/zap"
)

type Assets struct {
	client *rest.Client
}

func NewAssetsService(client *rest.Client) *Assets {
	return &Assets{
		client: client,
	}
}

func (s *Assets) GetAsset(assetId string) (*apiv1.Asset, error) {
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/assets/" + assetId).
		Do()
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump()))
		return nil, err
	}

	var asset apiv1.Asset
	if err := result.Into(&asset); err != nil {
		return nil, err
	}

	return &asset, nil
}

func (s *Assets) DeleteAsset(assetId string) error {
	result := s.client.
		Verb("DELETE").
		SetEndpoint("service/rest/v1/assets/" + assetId).
		Do()
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump()))
		return err
	}

	return nil
}

func (s *Assets) ListAssets(params *apiv1.ListAssetsParameters) (*apiv1.AssetList, error) {
	values, err := qencoder.Values(params)
	if err != nil {
		return nil, err
	}
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/assets").
		Param(values).
		Do()
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump()))
		return nil, err
	}

	var assets apiv1.AssetList
	if err := result.Into(&assets); err != nil {
		return nil, err
	}

	return &assets, nil
}
