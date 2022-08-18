package apiv1

import "time"

// Asset struct for Asset
type Asset struct {
	DownloadUrl    string            `json:"downloadUrl,omitempty"`
	Path           string            `json:"path,omitempty"`
	Id             string            `json:"id,omitempty"`
	Repository     string            `json:"repository,omitempty"`
	Format         string            `json:"format,omitempty"`
	Checksum       map[string]string `json:"checksum,omitempty"`
	ContentType    string            `json:"contentType,omitempty"`
	LastModified   time.Time         `json:"lastModified,omitempty"`
	LastDownloaded time.Time         `json:"lastDownloaded,omitempty"`
	FileSize       int64             `json:"fileSize,omitempty"`
}

// AssetList struct for AssetList
type AssetList struct {
	Items             []*Asset `json:"items,omitempty"`
	ContinuationToken string   `json:"continuationToken,omitempty"`
}
