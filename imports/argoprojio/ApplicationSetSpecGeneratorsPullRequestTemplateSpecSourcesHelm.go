package argoprojio


type ApplicationSetSpecGeneratorsPullRequestTemplateSpecSourcesHelm struct {
	ApiVersions *[]*string `field:"optional" json:"apiVersions" yaml:"apiVersions"`
	FileParameters *[]*ApplicationSetSpecGeneratorsPullRequestTemplateSpecSourcesHelmFileParameters `field:"optional" json:"fileParameters" yaml:"fileParameters"`
	IgnoreMissingValueFiles *bool `field:"optional" json:"ignoreMissingValueFiles" yaml:"ignoreMissingValueFiles"`
	KubeVersion *string `field:"optional" json:"kubeVersion" yaml:"kubeVersion"`
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	Parameters *[]*ApplicationSetSpecGeneratorsPullRequestTemplateSpecSourcesHelmParameters `field:"optional" json:"parameters" yaml:"parameters"`
	PassCredentials *bool `field:"optional" json:"passCredentials" yaml:"passCredentials"`
	ReleaseName *string `field:"optional" json:"releaseName" yaml:"releaseName"`
	SkipCrds *bool `field:"optional" json:"skipCrds" yaml:"skipCrds"`
	ValueFiles *[]*string `field:"optional" json:"valueFiles" yaml:"valueFiles"`
	Values *string `field:"optional" json:"values" yaml:"values"`
	ValuesObject interface{} `field:"optional" json:"valuesObject" yaml:"valuesObject"`
	Version *string `field:"optional" json:"version" yaml:"version"`
}

