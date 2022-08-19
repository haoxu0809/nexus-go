package components

import (
	"errors"

	"github.com/haoxu0809/nexus-go/components/apiv1"

	"github.com/haoxu0809/pkg/log"
	"github.com/haoxu0809/pkg/qencoder"
	"github.com/haoxu0809/pkg/rest"
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

func (s *Components) DeleteComponent(componentId string) error {
	result := s.client.
		Verb("DELETE").
		SetEndpoint("service/rest/v1/components/" + componentId).
		Do()
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump()))
		return err
	}

	return nil
}

func (s *Components) GetComponent(componentId string) (*apiv1.Component, error) {
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/components/" + componentId).
		Do()
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump()))
		return nil, err
	}

	var component apiv1.Component
	if err := result.Into(&component); err != nil {
		return nil, err
	}

	return &component, nil
}

func (s *Components) ListComponents(params *apiv1.ListComponentsParameters) (*apiv1.ComponentList, error) {
	values, err := qencoder.Values(params)
	if err != nil {
		return nil, err
	}
	result := s.client.
		Verb("GET").
		SetEndpoint("service/rest/v1/components").
		Param(values).
		Do()
	if err := result.Error(); err != nil {
		log.L().Error(err.Error(), zap.Object("context", result.Dump()))
		if result.StatusCode() == 404 {
			return nil, errors.New("404")
		}
		return nil, err
	}

	var components apiv1.ComponentList
	if err := result.Into(&components); err != nil {
		return nil, err
	}

	return &components, nil
}

func (s *Components) UploadComponent(r *apiv1.UploadComponentApiRequest) error {
	
	return nil
}
