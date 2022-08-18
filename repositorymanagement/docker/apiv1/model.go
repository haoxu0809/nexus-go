package apiv1

import (
	"github.com/haoxu0809/nexus-go/components/apiv1"
	"github.com/haoxu0809/nexus-go/repositorymanagement/common/attrs"
)

// DockerLocalRepository struct for DockerLocalRepository
type DockerLocalRepository struct {
	// A unique identifier for this repository
	Name string `json:"name,omitempty"`
	// Whether this repository accepts incoming requests
	Online    bool                 `json:"online"`
	Url       string               `json:"url"`
	Storage   *DockerLocalStorage  `json:"storage"`
	Cleanup   *attrs.CleanupPolicy `json:"cleanup,omitempty"`
	Component *apiv1.Attributes    `json:"component,omitempty"`
	Docker    *Attributes          `json:"docker"`
	Format    string               `json:"format"`
	Type      string               `json:"type"`
}

// DockerProxyRepository struct for DockerProxyRepository
type DockerProxyRepository struct {
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
	Docker          *Attributes        `json:"docker"`
	DockerProxy     *DockerProxy       `json:"dockerProxy"`
}
