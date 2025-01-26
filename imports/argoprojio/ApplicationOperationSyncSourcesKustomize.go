package argoprojio


// Kustomize holds kustomize specific options.
type ApplicationOperationSyncSourcesKustomize struct {
	// APIVersions specifies the Kubernetes resource API versions to pass to Helm when templating manifests.
	//
	// By default,
	// Argo CD uses the API versions of the target cluster. The format is [group/]version/kind.
	ApiVersions *[]*string `field:"optional" json:"apiVersions" yaml:"apiVersions"`
	// CommonAnnotations is a list of additional annotations to add to rendered manifests.
	CommonAnnotations *map[string]*string `field:"optional" json:"commonAnnotations" yaml:"commonAnnotations"`
	// CommonAnnotationsEnvsubst specifies whether to apply env variables substitution for annotation values.
	CommonAnnotationsEnvsubst *bool `field:"optional" json:"commonAnnotationsEnvsubst" yaml:"commonAnnotationsEnvsubst"`
	// CommonLabels is a list of additional labels to add to rendered manifests.
	CommonLabels *map[string]*string `field:"optional" json:"commonLabels" yaml:"commonLabels"`
	// Components specifies a list of kustomize components to add to the kustomization before building.
	Components *[]*string `field:"optional" json:"components" yaml:"components"`
	// ForceCommonAnnotations specifies whether to force applying common annotations to resources for Kustomize apps.
	ForceCommonAnnotations *bool `field:"optional" json:"forceCommonAnnotations" yaml:"forceCommonAnnotations"`
	// ForceCommonLabels specifies whether to force applying common labels to resources for Kustomize apps.
	ForceCommonLabels *bool `field:"optional" json:"forceCommonLabels" yaml:"forceCommonLabels"`
	// Images is a list of Kustomize image override specifications.
	Images *[]*string `field:"optional" json:"images" yaml:"images"`
	// KubeVersion specifies the Kubernetes API version to pass to Helm when templating manifests.
	//
	// By default, Argo CD
	// uses the Kubernetes version of the target cluster.
	KubeVersion *string `field:"optional" json:"kubeVersion" yaml:"kubeVersion"`
	// LabelWithoutSelector specifies whether to apply common labels to resource selectors or not.
	LabelWithoutSelector *bool `field:"optional" json:"labelWithoutSelector" yaml:"labelWithoutSelector"`
	// NamePrefix is a prefix appended to resources for Kustomize apps.
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// Namespace sets the namespace that Kustomize adds to all resources.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// NameSuffix is a suffix appended to resources for Kustomize apps.
	NameSuffix *string `field:"optional" json:"nameSuffix" yaml:"nameSuffix"`
	// Patches is a list of Kustomize patches.
	Patches *[]*ApplicationOperationSyncSourcesKustomizePatches `field:"optional" json:"patches" yaml:"patches"`
	// Replicas is a list of Kustomize Replicas override specifications.
	Replicas *[]*ApplicationOperationSyncSourcesKustomizeReplicas `field:"optional" json:"replicas" yaml:"replicas"`
	// Version controls which version of Kustomize to use for rendering manifests.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

