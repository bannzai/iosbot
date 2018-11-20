package hub

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/google/go-github/github"
)

type Service interface {
}

type ServiceImpl struct {
	Repository GitHubRepository
	Author     CommitAuthor
	Client     *github.Client
}

func NewService() Service {
	// TODO:
	return ServiceImpl{}
}

func (service ServiceImpl) DefaultBranch(ctx context.Context) (string, error) {
	repo, _, err := service.Client.Repositories.Get(ctx, service.Repository.Owner.Name, service.Repository.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch GitHub repository: %s", err)
	}
	defaultBranch := repo.GetDefaultBranch()
	return defaultBranch, nil
}

func (service *ServiceImpl) Branches(ctx context.Context) ([]*github.Branch, error) {
	branches, _, err := service.Client.Repositories.ListBranches(
		ctx,
		service.Repository.Owner,
		service.Repository.Name,
		&github.ListOptions{
			PerPage: 100,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch GitHub branches: %s", err)
	}
	return branches, nil
}

func (service *ServiceImpl) DownloadFile(ctx context.Context, sourceBranch, filePath string) ([]byte, error) {
	file, err := service.Client.Repositories.DownloadContents(
		ctx,
		service.Repository.Owner,
		service.Repository.Name,
		path,
		&github.RepositoryContentGetOptions{
			Ref: sourceBranch,
		},
	)

	defer func() {
		file.Close()
	}()

	if err != nil {
		return nil, fmt.Errorf("failed to download file from GitHub: %s", err)
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to download file from GitHub: %s", err)
	}

	return bytes, nil

}

func (service *ServiceImpl)
