package container_builder

import (
	"context"
	"fmt"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudbuild/v1"
)

// RepoNameFilterFromServiceName filters on the repo name, building the repo name using the service name
func RepoNameFilterWithServiceName(serviceName string) string {
	repoName := "github-vendasta-" + serviceName
	return fmt.Sprintf(`source.repo_source.repo_name="%s"`, repoName)
}

// BuiltImageFilter filters on the image id that resulted from the build
func BuiltImageFilter(imageID string) string {
	return fmt.Sprintf(`results.images.name="%s"`, imageID)
}

// List returns builds for the provided service
func List(ctx context.Context, filter string, pageSize int64, pageToken string) ([]*cloudbuild.Build, string, error) {
	client, err := google.DefaultClient(ctx, cloudbuild.CloudPlatformScope)
	if err != nil {
		return nil, "", err
	}
	cloudBuilder, err := cloudbuild.New(client)
	if err != nil {
		return nil, "", err
	}
	return (&ContainerBuilder{cloudBuilder}).List(ctx, filter, pageSize, pageToken)

}

type ContainerBuilder struct {
	cloudBuild *cloudbuild.Service
}

func (cb *ContainerBuilder) List(ctx context.Context, filter string, pageSize int64, pageToken string) ([]*cloudbuild.Build, string, error) {
	listCall := cb.cloudBuild.Projects.Builds.List("repcore-prod")
	listCall.Context(ctx)
	if filter != "" {
		listCall.Filter(filter)
	}
	if pageSize != 0 {
		listCall.PageSize(pageSize)
	}
	if pageToken != "" {
		listCall.PageToken(pageToken)
	}

	res, err := listCall.Do()
	if err != nil {
		return nil, "", err
	}
	return res.Builds, res.NextPageToken, nil
}
