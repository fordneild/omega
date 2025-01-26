package argoprojio


type ApplicationSetSpecGeneratorsMergeGeneratorsPluginTemplateSpecSyncPolicy struct {
	Automated *ApplicationSetSpecGeneratorsMergeGeneratorsPluginTemplateSpecSyncPolicyAutomated `field:"optional" json:"automated" yaml:"automated"`
	ManagedNamespaceMetadata *ApplicationSetSpecGeneratorsMergeGeneratorsPluginTemplateSpecSyncPolicyManagedNamespaceMetadata `field:"optional" json:"managedNamespaceMetadata" yaml:"managedNamespaceMetadata"`
	Retry *ApplicationSetSpecGeneratorsMergeGeneratorsPluginTemplateSpecSyncPolicyRetry `field:"optional" json:"retry" yaml:"retry"`
	SyncOptions *[]*string `field:"optional" json:"syncOptions" yaml:"syncOptions"`
}

