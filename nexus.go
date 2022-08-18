package nexus

import (
	"github.com/haoxu0809/nexus-go/assets"
	"github.com/haoxu0809/nexus-go/components"
	"github.com/haoxu0809/nexus-go/pkg/log"
	"github.com/haoxu0809/nexus-go/pkg/rest"
	"github.com/haoxu0809/nexus-go/repositorymanagement"
	"github.com/haoxu0809/nexus-go/search"
)

type Client struct {
	Assets               *assets.Assets
	Components           *components.Components
	RepositoryManagement *repositorymanagement.RepositoryManagement
	Search               *search.Search
}

func init() {
	log.New("debug")
}

func NewClient(cfg *rest.Config) (*Client, error) {
	c, err := rest.NewRESTClientForConfig(cfg)
	if err != nil {
		return nil, err
	}

	client := &Client{
		Assets:               assets.NewAssetsService(c),
		Components:           components.NewComponentsService(c),
		RepositoryManagement: repositorymanagement.NewRepositoryManagementService(c),
		Search:               search.NewSearchService(c),
	}

	return client, nil
}

func NewClientOrDie(cfg *rest.Config) *Client {
	client, err := NewClient(cfg)
	if err != nil {
		panic(err)
	}

	return client
}
