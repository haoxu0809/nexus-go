package components

import (
	"errors"

	"github.com/haoxu0809/nexus-go/components/apiv1"
	"github.com/haoxu0809/nexus-go/pkg/log"
	"github.com/haoxu0809/nexus-go/pkg/rest"
	"github.com/haoxu0809/nexus-go/pkg/xurl"

	"go.uber.org/zap"
)

type Components struct {
	client *rest.Client
}

func NewComponentsService(client *rest.Client) *Components {
	return &Components{
		client: client,
	}
}

func (s *Components) GetComponent(componentId string) (*apiv1.Component, error) {
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/components/" + componentId).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to get Component", zap.Object("context", log.Context(nil, result.Response())))
		return nil, result.Error()
	}

	var component apiv1.Component
	if err := result.Into(&component); err != nil {
		return nil, err
	}

	return &component, nil
}

func (s *Components) DeleteComponent(componentId string) error {
	result := s.client.
		Verb("DELETE").
		SetEndpoint("service/rest/v1/components/" + componentId).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to delete Component", zap.Object("context", log.Context(nil, result.Response())))
		return result.Error()
	}

	return nil
}

func (s *Components) ListComponents(params *apiv1.ListComponentsParameters) (*apiv1.ComponentList, error) {
	values, err := xurl.Values(params)
	if err != nil {
		return nil, err
	}
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/components").
		Param(values).
		Do()
	if result.Error() != nil {
		log.L().Error("failed to list Component", zap.Object("context", log.Context(nil, result.Response())))
		if result.Response().StatusCode == 404 {
			return nil, errors.New("404")
		}
		return nil, result.Error()
	}

	var components apiv1.ComponentList
	if err := result.Into(&components); err != nil {
		return nil, err
	}

	return &components, nil
}

func (s *Components) UploadComponent(component *apiv1.Component) error {
	// TODO
	return nil
}
