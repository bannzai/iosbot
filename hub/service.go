package hub

import (
	"context"
	"fmt"

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
