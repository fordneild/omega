package argoprojio


type ApplicationSetSpecGeneratorsPullRequestFilters struct {
	BranchMatch *string `field:"optional" json:"branchMatch" yaml:"branchMatch"`
	TargetBranchMatch *string `field:"optional" json:"targetBranchMatch" yaml:"targetBranchMatch"`
}

