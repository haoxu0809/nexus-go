package apiv1

type Parameters struct {
	Repository string `url:"repository,omitempty"`
	Name       string `url:"name,omitempty"`
	Version    string `url:"version,omitempty"`
	Sort       string `url:"sort,omitempty"`
	Direction  string `url:"direction,omitempty"`
}

type SearchAssetsParameters struct {
	ContinuationToken string `url:"continuationToken,omitempty"`
	*Parameters
}

type SearchAndDownloadAssetParameters struct {
	*Parameters
}

type SearchComponentsParameters struct {
	ContinuationToken string `url:"continuationToken,omitempty"`
	*Parameters
}
