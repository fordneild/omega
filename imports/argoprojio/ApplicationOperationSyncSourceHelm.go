package argoprojio


// Helm holds helm specific options.
type ApplicationOperationSyncSourceHelm struct {
	// APIVersions specifies the Kubernetes resource API versions to pass to Helm when templating manifests.
	//
	// By default,
	// Argo CD uses the API versions of the target cluster. The format is [group/]version/kind.
	ApiVersions *[]*string `field:"optional" json:"apiVersions" yaml:"apiVersions"`
	// FileParameters are file parameters to the helm template.
	FileParameters *[]*ApplicationOperationSyncSourceHelmFileParameters `field:"optional" json:"fileParameters" yaml:"fileParameters"`
	// IgnoreMissingValueFiles prevents helm template from failing when valueFiles do not exist locally by not appending them to helm template --values.
	IgnoreMissingValueFiles *bool `field:"optional" json:"ignoreMissingValueFiles" yaml:"ignoreMissingValueFiles"`
	// KubeVersion specifies the Kubernetes API version to pass to Helm when templating manifests.
	//
	// By default, Argo CD
	// uses the Kubernetes version of the target cluster.
	KubeVersion *string `field:"optional" json:"kubeVersion" yaml:"kubeVersion"`
	// Namespace is an optional namespace to template with.
	//
	// If left empty, defaults to the app's destination namespace.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Parameters is a list of Helm parameters which are passed to the helm template command upon manifest generation.
	Parameters *[]*ApplicationOperationSyncSourceHelmParameters `field:"optional" json:"parameters" yaml:"parameters"`
	// PassCredentials pass credentials to all domains (Helm's --pass-credentials).
	PassCredentials *bool `field:"optional" json:"passCredentials" yaml:"passCredentials"`
	// ReleaseName is the Helm release name to use.
	//
	// If omitted it will use the application name.
	ReleaseName *string `field:"optional" json:"releaseName" yaml:"releaseName"`
	// SkipCrds skips custom resource definition installation step (Helm's --skip-crds).
	SkipCrds *bool `field:"optional" json:"skipCrds" yaml:"skipCrds"`
	// ValuesFiles is a list of Helm value files to use when generating a template.
	ValueFiles *[]*string `field:"optional" json:"valueFiles" yaml:"valueFiles"`
	// Values specifies Helm values to be passed to helm template, typically defined as a block.
	//
	// ValuesObject takes precedence over Values, so use one or the other.
	Values *string `field:"optional" json:"values" yaml:"values"`
	// ValuesObject specifies Helm values to be passed to helm template, defined as a map.
	//
	// This takes precedence over Values.
	ValuesObject interface{} `field:"optional" json:"valuesObject" yaml:"valuesObject"`
	// Version is the Helm version to use for templating ("3").
	Version *string `field:"optional" json:"version" yaml:"version"`
}

