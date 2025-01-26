package argoprojio


type ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderAwsCodeCommit struct {
	AllBranches *bool `field:"optional" json:"allBranches" yaml:"allBranches"`
	Region *string `field:"optional" json:"region" yaml:"region"`
	Role *string `field:"optional" json:"role" yaml:"role"`
	TagFilters *[]*ApplicationSetSpecGeneratorsMergeGeneratorsScmProviderAwsCodeCommitTagFilters `field:"optional" json:"tagFilters" yaml:"tagFilters"`
}

