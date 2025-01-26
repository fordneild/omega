package argoprojio


type ApplicationSetSpecGeneratorsMatrixGeneratorsPullRequestBitbucket struct {
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	Repo *string `field:"required" json:"repo" yaml:"repo"`
	Api *string `field:"optional" json:"api" yaml:"api"`
	BasicAuth *ApplicationSetSpecGeneratorsMatrixGeneratorsPullRequestBitbucketBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	BearerToken *ApplicationSetSpecGeneratorsMatrixGeneratorsPullRequestBitbucketBearerToken `field:"optional" json:"bearerToken" yaml:"bearerToken"`
}

