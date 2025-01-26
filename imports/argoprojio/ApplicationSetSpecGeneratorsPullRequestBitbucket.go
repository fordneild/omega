package argoprojio


type ApplicationSetSpecGeneratorsPullRequestBitbucket struct {
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	Repo *string `field:"required" json:"repo" yaml:"repo"`
	Api *string `field:"optional" json:"api" yaml:"api"`
	BasicAuth *ApplicationSetSpecGeneratorsPullRequestBitbucketBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	BearerToken *ApplicationSetSpecGeneratorsPullRequestBitbucketBearerToken `field:"optional" json:"bearerToken" yaml:"bearerToken"`
}

