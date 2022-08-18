package repositorymanagement

import (
	"github.com/haoxu0809/nexus-go/pkg/rest"
	"github.com/haoxu0809/nexus-go/repositorymanagement/cocoapods"
	"github.com/haoxu0809/nexus-go/repositorymanagement/docker"
	"github.com/haoxu0809/nexus-go/repositorymanagement/gitlfs"
	"github.com/haoxu0809/nexus-go/repositorymanagement/golang"
	"github.com/haoxu0809/nexus-go/repositorymanagement/helm"
	"github.com/haoxu0809/nexus-go/repositorymanagement/maven"
	"github.com/haoxu0809/nexus-go/repositorymanagement/npm"
	"github.com/haoxu0809/nexus-go/repositorymanagement/pypi"
	"github.com/haoxu0809/nexus-go/repositorymanagement/raw"
	"github.com/haoxu0809/nexus-go/repositorymanagement/repositories"
)

type RepositoryManagement struct {
	*dockerrepo.Docker
	*gitlfsrepo.GitLfs
	*golangrepo.Golang
	*helmrepo.Helm
	*mavenrepo.Maven
	*npmrepo.Npm
	*pypirepo.Pypi
	*rawrepo.Raw
	*cocoapodsrepo.CocoaPods
	*repositories.Repository
}

func NewRepositoryManagementService(client *rest.Client) *RepositoryManagement {
	return &RepositoryManagement{
		Docker:     dockerrepo.NewDockerService(client),
		GitLfs:     gitlfsrepo.NewGitLfsService(client),
		Golang:     golangrepo.NewGolangService(client),
		Helm:       helmrepo.NewHelmService(client),
		Maven:      mavenrepo.NewMavenService(client),
		Npm:        npmrepo.NewNpmService(client),
		Pypi:       pypirepo.NewPypiService(client),
		Raw:        rawrepo.NewRawService(client),
		CocoaPods:  cocoapodsrepo.NewCocoaPodsService(client),
		Repository: repositories.NewRepositoryService(client),
	}
}
