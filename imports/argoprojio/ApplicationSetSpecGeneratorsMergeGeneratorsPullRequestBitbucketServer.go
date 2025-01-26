package argoprojio


type ApplicationSetSpecGeneratorsMergeGeneratorsPullRequestBitbucketServer struct {
	Api *string `field:"required" json:"api" yaml:"api"`
	Project *string `field:"required" json:"project" yaml:"project"`
	Repo *string `field:"required" json:"repo" yaml:"repo"`
	BasicAuth *ApplicationSetSpecGeneratorsMergeGeneratorsPullRequestBitbucketServerBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	BearerToken *ApplicationSetSpecGeneratorsMergeGeneratorsPullRequestBitbucketServerBearerToken `field:"optional" json:"bearerToken" yaml:"bearerToken"`
	CaRef *ApplicationSetSpecGeneratorsMergeGeneratorsPullRequestBitbucketServerCaRef `field:"optional" json:"caRef" yaml:"caRef"`
	Insecure *bool `field:"optional" json:"insecure" yaml:"insecure"`
}

