package apiv1

// Repository struct for Repository
type Repository struct {
	Name       string                            `json:"name,omitempty"`
	Format     string                            `json:"format,omitempty"`
	Type       string                            `json:"type,omitempty"`
	Url        string                            `json:"url,omitempty"`
	Attributes map[string]map[string]interface{} `json:"attributes,omitempty"`
}

// RepositorySettings struct for RepositorySettings
type RepositorySettings struct {
	// A unique identifier for this repository
	Name string `json:"name,omitempty"`
	// Component format held in this repository
	Format string `json:"format,omitempty"`
	// Controls if deployments of and updates to artifacts are allowed
	Type string `json:"type,omitempty"`
	// URL to the repository
	Url string `json:"url,omitempty"`
	// Whether this repository accepts incoming requests
	Online bool `json:"online"`
}
