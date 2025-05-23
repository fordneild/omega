package argoprojio


type ApplicationSetSpecGeneratorsMergeGeneratorsPullRequest struct {
	Azuredevops *ApplicationSetSpecGeneratorsMergeGeneratorsPullRequestAzuredevops `field:"optional" json:"azuredevops" yaml:"azuredevops"`
	Bitbucket *ApplicationSetSpecGeneratorsMergeGeneratorsPullRequestBitbucket `field:"optional" json:"bitbucket" yaml:"bitbucket"`
	BitbucketServer *ApplicationSetSpecGeneratorsMergeGeneratorsPullRequestBitbucketServer `field:"optional" json:"bitbucketServer" yaml:"bitbucketServer"`
	Filters *[]*ApplicationSetSpecGeneratorsMergeGeneratorsPullRequestFilters `field:"optional" json:"filters" yaml:"filters"`
	Gitea *ApplicationSetSpecGeneratorsMergeGeneratorsPullRequestGitea `field:"optional" json:"gitea" yaml:"gitea"`
	Github *ApplicationSetSpecGeneratorsMergeGeneratorsPullRequestGithub `field:"optional" json:"github" yaml:"github"`
	Gitlab *ApplicationSetSpecGeneratorsMergeGeneratorsPullRequestGitlab `field:"optional" json:"gitlab" yaml:"gitlab"`
	RequeueAfterSeconds *float64 `field:"optional" json:"requeueAfterSeconds" yaml:"requeueAfterSeconds"`
	Template *ApplicationSetSpecGeneratorsMergeGeneratorsPullRequestTemplate `field:"optional" json:"template" yaml:"template"`
}

