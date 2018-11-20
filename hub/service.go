package hub

import "github.com/google/go-github/github"

type Service interface {
}

type ServiceImpl struct {
	Repository GitHubRepository
	Author     CommitAuthor
	Client     *github.Client
}
