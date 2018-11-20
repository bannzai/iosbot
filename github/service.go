package github

type Service interface {
}

type ServiceImpl struct {
	Repository GitHubRepository
	Author     CommitAuthor
	Client     *github.Client
}
