package apiv1

import (
	"github.com/haoxu0809/nexus-go/assets/apiv1"
)

// Attributes struct for Attributes
type Attributes struct {
	// Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	ProprietaryComponents *bool `json:"proprietaryComponents,omitempty"`
}

// Component struct for Component
type Component struct {
	Id         *string        `json:"id,omitempty"`
	Repository *string        `json:"repository,omitempty"`
	Format     *string        `json:"format,omitempty"`
	Group      *string        `json:"group,omitempty"`
	Name       *string        `json:"name,omitempty"`
	Version    *string        `json:"version,omitempty"`
	Assets     []*apiv1.Asset `json:"assets,omitempty"`
}

// ComponentList struct for ComponentList
type ComponentList struct {
	Items             []*Component `json:"items,omitempty"`
	ContinuationToken *string      `json:"continuationToken,omitempty"`
}
