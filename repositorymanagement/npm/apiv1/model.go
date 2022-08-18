package apiv1

import (
	"github.com/haoxu0809/nexus-go/repositorymanagement/common/attrs"
)

// NpmProxyRepository struct for NpmProxyRepository
type NpmProxyRepository struct {
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
	Npm             *Attributes        `json:"npm,omitempty"`
}
