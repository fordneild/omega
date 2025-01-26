package argoprojio


type ApplicationSetSpecGeneratorsMergeGeneratorsClusterDecisionResourceTemplateSpecSourcesKustomize struct {
	ApiVersions *[]*string `field:"optional" json:"apiVersions" yaml:"apiVersions"`
	CommonAnnotations *map[string]*string `field:"optional" json:"commonAnnotations" yaml:"commonAnnotations"`
	CommonAnnotationsEnvsubst *bool `field:"optional" json:"commonAnnotationsEnvsubst" yaml:"commonAnnotationsEnvsubst"`
	CommonLabels *map[string]*string `field:"optional" json:"commonLabels" yaml:"commonLabels"`
	Components *[]*string `field:"optional" json:"components" yaml:"components"`
	ForceCommonAnnotations *bool `field:"optional" json:"forceCommonAnnotations" yaml:"forceCommonAnnotations"`
	ForceCommonLabels *bool `field:"optional" json:"forceCommonLabels" yaml:"forceCommonLabels"`
	Images *[]*string `field:"optional" json:"images" yaml:"images"`
	KubeVersion *string `field:"optional" json:"kubeVersion" yaml:"kubeVersion"`
	LabelWithoutSelector *bool `field:"optional" json:"labelWithoutSelector" yaml:"labelWithoutSelector"`
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	NameSuffix *string `field:"optional" json:"nameSuffix" yaml:"nameSuffix"`
	Patches *[]*ApplicationSetSpecGeneratorsMergeGeneratorsClusterDecisionResourceTemplateSpecSourcesKustomizePatches `field:"optional" json:"patches" yaml:"patches"`
	Replicas *[]*ApplicationSetSpecGeneratorsMergeGeneratorsClusterDecisionResourceTemplateSpecSourcesKustomizeReplicas `field:"optional" json:"replicas" yaml:"replicas"`
	Version *string `field:"optional" json:"version" yaml:"version"`
}

