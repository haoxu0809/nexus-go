package apiv1

import "os"

type ListComponentsParameters struct {
	Repository        string `url:"repository"`
	ContinuationToken string `url:"continuationToken,omitempty"`
}

type UploadComponentApiRequest struct {
	Repository             string  `form:"repository"`
	RAsset                 os.File `form:"r.asset"`
	RAssetPathId           string  `form:"r.asset.pathId"`
	PypiAsset              os.File `form:"pypi.asset"`
	HelmAsset              os.File `form:"helm.asset"`
	YumDirectory           string  `form:"yum.directory"`
	YumAsset               os.File `form:"yum.asset"`
	YumAssetFilename       string  `form:"yum.asset.filename"`
	DockerAsset            os.File `form:"docker.asset"`
	RubygemsAsset          os.File `form:"rubygems.asset"`
	NugetAsset             os.File `form:"nuget.asset"`
	NpmAsset               os.File `form:"npm.asset"`
	RawDirectory           string  `form:"raw.directory"`
	RawAsset1              os.File `form:"raw.asset1"`
	RawAsset1Filename      string  `form:"raw.asset1.filename"`
	AptAsset               os.File `form:"apt.asset"`
	Maven2GroupId          string  `form:"maven2.groupId"`
	Maven2ArtifactId       string  `form:"maven2.artifactId"`
	Maven2Version          string  `form:"maven2.version"`
	Maven2GeneratePom      bool    `form:"maven2.generate-pom"`
	Maven2Packaging        string  `form:"maven2.packaging"`
	Maven2Asset1           os.File `form:"maven2.asset1"`
	Maven2Asset1Classifier string  `form:"maven2.asset1.classifier"`
	Maven2Asset1Extension  string  `form:"maven2.asset1.extension"`
}
