package apiv1

type ListAssetsParameters struct {
	Repository        string `url:"repository"`
	ContinuationToken string `url:"continuationToken,omitempty"`
}
