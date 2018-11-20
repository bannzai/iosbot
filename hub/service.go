package hub

import git "github.com/google/go-github"

type Service interface {
}

type ServiceImpl struct {
	Repository GitHubRepository
	Author     CommitAuthor
	Client     *git.Client
}
