package argoprojio


type ApplicationSetSpecGeneratorsScmProviderGitlab struct {
	Group *string `field:"required" json:"group" yaml:"group"`
	AllBranches *bool `field:"optional" json:"allBranches" yaml:"allBranches"`
	Api *string `field:"optional" json:"api" yaml:"api"`
	CaRef *ApplicationSetSpecGeneratorsScmProviderGitlabCaRef `field:"optional" json:"caRef" yaml:"caRef"`
	IncludeSharedProjects *bool `field:"optional" json:"includeSharedProjects" yaml:"includeSharedProjects"`
	IncludeSubgroups *bool `field:"optional" json:"includeSubgroups" yaml:"includeSubgroups"`
	Insecure *bool `field:"optional" json:"insecure" yaml:"insecure"`
	TokenRef *ApplicationSetSpecGeneratorsScmProviderGitlabTokenRef `field:"optional" json:"tokenRef" yaml:"tokenRef"`
	Topic *string `field:"optional" json:"topic" yaml:"topic"`
}

