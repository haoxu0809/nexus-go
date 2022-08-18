package assets

import (
	"github.com/haoxu0809/nexus-go/assets/apiv1"
	"github.com/haoxu0809/nexus-go/pkg/log"
	"github.com/haoxu0809/nexus-go/pkg/rest"
	"github.com/haoxu0809/nexus-go/pkg/xurl"

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
	if result.Error() != nil {
		log.L().Error("failed to get Asset", zap.Object("context", log.Context(nil, result.Response())))
		return nil, result.Error()
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
	if result.Error() != nil {
		log.L().Error("failed to delete Asset", zap.Object("context", log.Context(nil, result.Response())))
		return result.Error()
	}

	return nil
}

func (s *Assets) ListAssets(params *apiv1.ListAssetsParameters) (*apiv1.AssetList, error) {
	values, err := xurl.Values(params)
	if err != nil {
		return nil, err
	}
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/assets").
		Param(values).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to list Asset", zap.Object("context", log.Context(nil, result.Response())))
		return nil, result.Error()
	}

	var assets apiv1.AssetList
	if err := result.Into(&assets); err != nil {
		return nil, err
	}

	return &assets, nil
}
