package apiv1

import (
	componentsapiv1 "github.com/haoxu0809/nexus-go/components/apiv1"
	"github.com/haoxu0809/nexus-go/repositorymanagement/common/attrs"
)

// GitLfsLocalApiRequest struct for GitLfsLocalApiRequest
type GitLfsLocalApiRequest struct {
	// A unique identifier for this repository
	Name string `json:"name"`
	// Whether this repository accepts incoming requests
	Online    bool                        `json:"online"`
	Storage   *attrs.LocalStorage         `json:"storage"`
	Cleanup   *attrs.CleanupPolicy        `json:"cleanup,omitempty"`
	Component *componentsapiv1.Attributes `json:"component,omitempty"`
}
