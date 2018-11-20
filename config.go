package main

type Config struct {
	BotToken              string `toml:"bot_token"`
	VerificationToken     string `toml:"verification_token"`
	BotID                 string `toml:"bot_id"`
	ChannelID             string `toml:"channel_id"`
	DebugChannelID        string `toml:"debug_channel_id"`
	GitHubUsername        string `toml:"github_username"`
	GitHubToken           string `toml:"github_token"`
	GitHubRepositoryOwner string `toml:"github_repository_owner"`
	GitHubRepositoryName  string `toml:"github_repository_name"`
	GitCommitAuthorName   string `toml:"github_commit_author_name"`
	GitCommitAuthorEmail  string `toml:"github_commit_author_email"`
	InfoPlistPath         string `toml:"infoplist_path"`
}
