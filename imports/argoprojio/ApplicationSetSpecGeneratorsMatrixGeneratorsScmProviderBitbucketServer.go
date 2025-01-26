package argoprojio


type ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderBitbucketServer struct {
	Api *string `field:"required" json:"api" yaml:"api"`
	Project *string `field:"required" json:"project" yaml:"project"`
	AllBranches *bool `field:"optional" json:"allBranches" yaml:"allBranches"`
	BasicAuth *ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderBitbucketServerBasicAuth `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	BearerToken *ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderBitbucketServerBearerToken `field:"optional" json:"bearerToken" yaml:"bearerToken"`
	CaRef *ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderBitbucketServerCaRef `field:"optional" json:"caRef" yaml:"caRef"`
	Insecure *bool `field:"optional" json:"insecure" yaml:"insecure"`
}

