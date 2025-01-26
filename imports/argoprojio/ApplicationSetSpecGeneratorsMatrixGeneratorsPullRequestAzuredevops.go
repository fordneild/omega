package argoprojio


type ApplicationSetSpecGeneratorsMatrixGeneratorsPullRequestAzuredevops struct {
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	Project *string `field:"required" json:"project" yaml:"project"`
	Repo *string `field:"required" json:"repo" yaml:"repo"`
	Api *string `field:"optional" json:"api" yaml:"api"`
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	TokenRef *ApplicationSetSpecGeneratorsMatrixGeneratorsPullRequestAzuredevopsTokenRef `field:"optional" json:"tokenRef" yaml:"tokenRef"`
}

