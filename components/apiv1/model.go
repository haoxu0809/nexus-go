package apiv1

import (
	"github.com/haoxu0809/nexus-go/assets/apiv1"
)

// Attributes struct for Attributes
type Attributes struct {
	// Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)
	ProprietaryComponents bool `json:"proprietaryComponents"`
}

// Component struct for Component
type Component struct {
	Id         string        `json:"id"`
	Repository string        `json:"repository"`
	Format     string        `json:"format"`
	Group      string        `json:"group"`
	Name       string        `json:"name"`
	Version    string        `json:"version"`
	Assets     []apiv1.Asset `json:"assets"`
}

// ComponentList struct for ComponentList
type ComponentList struct {
	Items             []Component `json:"items"`
	ContinuationToken string      `json:"continuationToken"`
}
