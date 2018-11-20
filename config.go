package main

type Config struct {
	BotToken              string `requiered:"true" envconfig:"BOT_TOKEN"`
	BotID                 string `requiered:"true" envconfig:"BOT_ID"`
	ChannelID             string `requiered:"true" envconfig:"CHANNEL_ID"`
	DebugChannelID        string `requiered:"true" envconfig:"DEBUG_CHANNEL_ID"`
	GitHubUsername        string `requiered:"true" envconfig:"GITHUB_USERNAME"`
	GitHubToken           string `requiered:"true" envconfig:"GITHUB_TOKEN"`
	GitHubRepositoryOwner string `requiered:"true" envconfig:"GITHUB_REPOSITORY_OWNER"`
	GitHubRepositoryName  string `requiered:"true" envconfig:"GITHUB_REPOSITORY_NAME"`
	GitCommitAuthorName   string `requiered:"true" envconfig:"GIT_COMMIT_AUTHOR_NAME"`
	GitCommitAuthorEmail  string `requiered:"true" envconfig:"GIT_COMMIT_AUTHOR_EMAIL"`
	InfoPlistPath         string `requiered:"true" envconfig:"INFOPLIST_PATH"`
}
