package argoprojio


type ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderTemplateSpecSyncPolicy struct {
	Automated *ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderTemplateSpecSyncPolicyAutomated `field:"optional" json:"automated" yaml:"automated"`
	ManagedNamespaceMetadata *ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderTemplateSpecSyncPolicyManagedNamespaceMetadata `field:"optional" json:"managedNamespaceMetadata" yaml:"managedNamespaceMetadata"`
	Retry *ApplicationSetSpecGeneratorsMatrixGeneratorsScmProviderTemplateSpecSyncPolicyRetry `field:"optional" json:"retry" yaml:"retry"`
	SyncOptions *[]*string `field:"optional" json:"syncOptions" yaml:"syncOptions"`
}

