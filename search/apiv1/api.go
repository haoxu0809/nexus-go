package apiv1

type Query struct {
	Repository string `url:"repository,omitempty"`
	Name       string `url:"name,omitempty"`
	Version    string `url:"version,omitempty"`
	Sort       string `url:"sort,omitempty"`
	Direction  string `url:"direction,omitempty"`
}

type Assets struct {
	Next string `url:"continuationToken,omitempty"`
	Query
}

type DownloadAsset struct {
	Query
}

type Components struct {
	ContinuationToken string `url:"continuationToken,omitempty"`
	Query
}
