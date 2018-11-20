package hub

type PullRequest struct {
	TargetBranch  string
	CommitBranch  string
	FileContent   []byte
	FilePath      string
	Title         string
	CommitMessage string
}
