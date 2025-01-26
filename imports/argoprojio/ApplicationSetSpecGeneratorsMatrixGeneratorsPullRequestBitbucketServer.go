package argoprojio


type ApplicationSetSpecGeneratorsMatrixGeneratorsPullRequestBitbucketServer struct {
	Api *string `field:"required" json:"api" yaml:"api"`
	Project *string `field:"required" json:"project" yaml:"project"`
	Repo *string `field:"required" json:"repo" yaml:"repo"`
	BasicAuth *ApplicationSetSpecGeneratorsMatrixGeneratorsPullRequestBitbucketServerBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	BearerToken *ApplicationSetSpecGeneratorsMatrixGeneratorsPullRequestBitbucketServerBearerToken `field:"optional" json:"bearerToken" yaml:"bearerToken"`
	CaRef *ApplicationSetSpecGeneratorsMatrixGeneratorsPullRequestBitbucketServerCaRef `field:"optional" json:"caRef" yaml:"caRef"`
	Insecure *bool `field:"optional" json:"insecure" yaml:"insecure"`
}

