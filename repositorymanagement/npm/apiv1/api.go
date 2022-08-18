package apiv1

import (
	"github.com/haoxu0809/nexus-go/components/apiv1"
	"github.com/haoxu0809/nexus-go/repositorymanagement/common/attrs"
)

// Attributes struct for Attributes
type Attributes struct {
	// Remove Non-Cataloged Versions
	RemoveNonCataloged bool `json:"removeNonCataloged"`
	// Remove Quarantined Versions
	RemoveQuarantined bool `json:"removeQuarantined"`
}

// NpmLocalApiRequest struct for NpmLocalApiRequest
type NpmLocalApiRequest struct {
	// A unique identifier for this repository
	Name string `json:"name"`
	// Whether this repository accepts incoming requests
	Online    bool                 `json:"online"`
	Storage   *attrs.LocalStorage  `json:"storage"`
	Cleanup   *attrs.CleanupPolicy `json:"cleanup,omitempty"`
	Component *apiv1.Attributes    `json:"component,omitempty"`
}
