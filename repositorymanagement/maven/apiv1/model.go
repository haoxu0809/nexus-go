package apiv1

import (
	componentsapiv1 "github.com/haoxu0809/nexus-go/components/apiv1"
	"github.com/haoxu0809/nexus-go/repositorymanagement/common/attrs"
)

// MavenLocalRepository struct for MavenLocalRepository
type MavenLocalRepository struct {
	// A unique identifier for this repository
	Name string `json:"name,omitempty"`
	// Whether this repository accepts incoming requests
	Online    bool                        `json:"online"`
	Storage   *attrs.LocalStorage         `json:"storage"`
	Cleanup   *attrs.CleanupPolicy        `json:"cleanup,omitempty"`
	Component *componentsapiv1.Attributes `json:"component,omitempty"`
	Maven     *Attributes                 `json:"maven"`
}

// MavenProxyRepository struct for MavenProxyRepository
type MavenProxyRepository struct {
	// A unique identifier for this repository
	Name string `json:"name,omitempty"`
	// Whether this repository accepts incoming requests
	Online        bool                 `json:"online"`
	Storage       *attrs.Storage       `json:"storage"`
	Cleanup       *attrs.CleanupPolicy `json:"cleanup,omitempty"`
	Proxy         *attrs.Proxy         `json:"proxy"`
	NegativeCache *attrs.NegativeCache `json:"negativeCache"`
	HttpClient    *attrs.HttpClient    `json:"httpClient"`
	// The name of the routing rule assigned to this repository
	RoutingRuleName string             `json:"routingRuleName,omitempty"`
	Replication     *attrs.Replication `json:"replication,omitempty"`
	Maven           *Attributes        `json:"maven"`
}
