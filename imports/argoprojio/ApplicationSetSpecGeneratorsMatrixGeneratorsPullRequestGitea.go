package argoprojio


type ApplicationSetSpecGeneratorsMatrixGeneratorsPullRequestGitea struct {
	Api *string `field:"required" json:"api" yaml:"api"`
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	Repo *string `field:"required" json:"repo" yaml:"repo"`
	Insecure *bool `field:"optional" json:"insecure" yaml:"insecure"`
	TokenRef *ApplicationSetSpecGeneratorsMatrixGeneratorsPullRequestGiteaTokenRef `field:"optional" json:"tokenRef" yaml:"tokenRef"`
}

