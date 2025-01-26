package argoprojio


type ApplicationSetSpecGeneratorsMergeGeneratorsPullRequestBitbucket struct {
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	Repo *string `field:"required" json:"repo" yaml:"repo"`
	Api *string `field:"optional" json:"api" yaml:"api"`
	BasicAuth *ApplicationSetSpecGeneratorsMergeGeneratorsPullRequestBitbucketBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	BearerToken *ApplicationSetSpecGeneratorsMergeGeneratorsPullRequestBitbucketBearerToken `field:"optional" json:"bearerToken" yaml:"bearerToken"`
}

