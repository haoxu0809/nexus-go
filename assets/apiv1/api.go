package apiv1

type ListAssetsParameters struct {
	Repository        string `form:"repository,required"`
	ContinuationToken string `form:"continuationToken,omitempty"`
}
