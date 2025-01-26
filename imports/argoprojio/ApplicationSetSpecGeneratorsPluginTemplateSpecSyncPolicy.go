package argoprojio


type ApplicationSetSpecGeneratorsPluginTemplateSpecSyncPolicy struct {
	Automated *ApplicationSetSpecGeneratorsPluginTemplateSpecSyncPolicyAutomated `field:"optional" json:"automated" yaml:"automated"`
	ManagedNamespaceMetadata *ApplicationSetSpecGeneratorsPluginTemplateSpecSyncPolicyManagedNamespaceMetadata `field:"optional" json:"managedNamespaceMetadata" yaml:"managedNamespaceMetadata"`
	Retry *ApplicationSetSpecGeneratorsPluginTemplateSpecSyncPolicyRetry `field:"optional" json:"retry" yaml:"retry"`
	SyncOptions *[]*string `field:"optional" json:"syncOptions" yaml:"syncOptions"`
}

