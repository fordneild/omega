package argoprojio


type ApplicationSetSpecGeneratorsScmProviderAwsCodeCommit struct {
	AllBranches *bool `field:"optional" json:"allBranches" yaml:"allBranches"`
	Region *string `field:"optional" json:"region" yaml:"region"`
	Role *string `field:"optional" json:"role" yaml:"role"`
	TagFilters *[]*ApplicationSetSpecGeneratorsScmProviderAwsCodeCommitTagFilters `field:"optional" json:"tagFilters" yaml:"tagFilters"`
}

