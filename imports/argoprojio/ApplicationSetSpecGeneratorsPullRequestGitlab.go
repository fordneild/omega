package argoprojio


type ApplicationSetSpecGeneratorsPullRequestGitlab struct {
	Project *string `field:"required" json:"project" yaml:"project"`
	Api *string `field:"optional" json:"api" yaml:"api"`
	CaRef *ApplicationSetSpecGeneratorsPullRequestGitlabCaRef `field:"optional" json:"caRef" yaml:"caRef"`
	Insecure *bool `field:"optional" json:"insecure" yaml:"insecure"`
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	PullRequestState *string `field:"optional" json:"pullRequestState" yaml:"pullRequestState"`
	TokenRef *ApplicationSetSpecGeneratorsPullRequestGitlabTokenRef `field:"optional" json:"tokenRef" yaml:"tokenRef"`
}

