package main

type Config struct {
	BotToken              string `envconfig:"BOT_TOKEN"`
	BotID                 string `envconfig:"BOT_ID"`
	ChannelID             string `envconfig:"CHANNEL_ID"`
	DebugChannelID        string `envconfig:"DEBUG_CHANNEL_ID"`
	GitHubUsername        string `envconfig:"GITHUB_USERNAME"`
	GitHubToken           string `envconfig:"GITHUB_TOKEN"`
	GitHubRepositoryOwner string `envconfig:"GITHUB_REPOSITORY_OWNER"`
	GitHubRepositoryName  string `envconfig:"GITHUB_REPOSITORY_NAME"`
	GitCommitAuthorName   string `envconfig:"GIT_COMMIT_AUTHOR_NAME"`
	GitCommitAuthorEmail  string `envconfig:"GIT_COMMIT_AUTHOR_EMAIL"`
	InfoPlistPath         string `envconfig:"INFOPLIST_PATH"`
}
