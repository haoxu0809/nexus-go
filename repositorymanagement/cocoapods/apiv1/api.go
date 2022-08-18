package apiv1

import "github.com/haoxu0809/nexus-go/repositorymanagement/common/attrs"

// CocoapodsProxyApiRequest struct for CocoapodsProxyApiRequest
type CocoapodsProxyApiRequest struct {
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
