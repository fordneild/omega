package argoprojio


type ApplicationSetSpecGeneratorsMatrixGeneratorsPluginTemplateSpecSyncPolicy struct {
	Automated *ApplicationSetSpecGeneratorsMatrixGeneratorsPluginTemplateSpecSyncPolicyAutomated `field:"optional" json:"automated" yaml:"automated"`
	ManagedNamespaceMetadata *ApplicationSetSpecGeneratorsMatrixGeneratorsPluginTemplateSpecSyncPolicyManagedNamespaceMetadata `field:"optional" json:"managedNamespaceMetadata" yaml:"managedNamespaceMetadata"`
	Retry *ApplicationSetSpecGeneratorsMatrixGeneratorsPluginTemplateSpecSyncPolicyRetry `field:"optional" json:"retry" yaml:"retry"`
	SyncOptions *[]*string `field:"optional" json:"syncOptions" yaml:"syncOptions"`
}

