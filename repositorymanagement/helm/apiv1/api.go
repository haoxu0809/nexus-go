package apiv1

import (
	componentsapiv1 "github.com/haoxu0809/nexus-go/components/apiv1"
	"github.com/haoxu0809/nexus-go/repositorymanagement/common/attrs"
)

// HelmLocalApiRequest struct for HelmLocalApiRequest
type HelmLocalApiRequest struct {
	// A unique identifier for this repository
	Name string `json:"name"`
	// Whether this repository accepts incoming requests
	Online    bool                        `json:"online"`
	Storage   *attrs.LocalStorage         `json:"storage"`
	Cleanup   *attrs.CleanupPolicy        `json:"cleanup,omitempty"`
	Component *componentsapiv1.Attributes `json:"component,omitempty"`
}

// HelmProxyApiRequest struct for HelmProxyApiRequest
type HelmProxyApiRequest struct {
	// A unique identifier for this repository
	Name string `json:"name"`
	// Whether this repository accepts incoming requests
	Online        bool                 `json:"online"`
	Storage       *attrs.Storage       `json:"storage"`
	Cleanup       *attrs.CleanupPolicy `json:"cleanup,omitempty"`
	Proxy         *attrs.Proxy         `json:"proxy"`
	NegativeCache *attrs.NegativeCache `json:"negativeCache"`
	HttpClient    *attrs.HttpClient    `json:"httpClient"`
	RoutingRule   string               `json:"routingRule,omitempty"`
	Replication   *attrs.Replication   `json:"replication,omitempty"`
}
