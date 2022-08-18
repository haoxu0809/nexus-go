package apiv1

import (
	componentsapiv1 "github.com/haoxu0809/nexus-go/components/apiv1"
	"github.com/haoxu0809/nexus-go/repositorymanagement/common/attrs"
)

// Attributes struct for Attributes
type Attributes struct {
	// Whether to allow clients to use the V1 API to interact with this repository
	V1Enabled bool `json:"v1Enabled"`
	// Whether to force authentication (Docker Bearer Token Realm required if false)
	ForceBasicAuth bool `json:"forceBasicAuth"`
	// Create an HTTP connector at specified port
	HttpPort int32 `json:"httpPort,omitempty"`
	// Create an HTTPS connector at specified port
	HttpsPort int32 `json:"httpsPort,omitempty"`
	// Allows to use subdomain
	Subdomain string `json:"subdomain,omitempty"`
}

// DockerLocalStorage struct for DockerLocalStorage
type DockerLocalStorage struct {
	// Blob store used to store repository contents
	BlobStoreName string `json:"blobStoreName,omitempty"`
	// Whether to validate uploaded content's MIME type appropriate for the repository format
	StrictContentTypeValidation bool `json:"strictContentTypeValidation"`
	// Controls if deployments of and updates to assets are allowed
	WritePolicy string `json:"writePolicy"`
	// Whether to allow redeploying the 'latest' tag but defer to the Deployment Policy for all other tags
	LatestPolicy bool `json:"latestPolicy,omitempty"`
}

// DockerProxy struct for DockerProxy
type DockerProxy struct {
	// Type of Docker Index
	IndexType string `json:"indexType,omitempty"`
	// Url of Docker Index to use
	IndexUrl string `json:"indexUrl,omitempty"`
}

// DockerLocalApiRequest struct for DockerLocalApiRequest
type DockerLocalApiRequest struct {
	// A unique identifier for this repository
	Name string `json:"name"`
	// Whether this repository accepts incoming requests
	Online    bool                        `json:"online"`
	Storage   *DockerLocalStorage         `json:"storage"`
	Cleanup   *attrs.CleanupPolicy        `json:"cleanup,omitempty"`
	Component *componentsapiv1.Attributes `json:"component,omitempty"`
	Docker    *Attributes                 `json:"docker"`
}

// DockerProxyApiRequest struct for DockerProxyApiRequest
type DockerProxyApiRequest struct {
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
	Docker        *Attributes          `json:"docker"`
	DockerProxy   *DockerProxy         `json:"dockerProxy"`
}
