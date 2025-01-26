package argoprojio


type ApplicationSetSpecGeneratorsMatrixGeneratorsListTemplateSpecSyncPolicy struct {
	Automated *ApplicationSetSpecGeneratorsMatrixGeneratorsListTemplateSpecSyncPolicyAutomated `field:"optional" json:"automated" yaml:"automated"`
	ManagedNamespaceMetadata *ApplicationSetSpecGeneratorsMatrixGeneratorsListTemplateSpecSyncPolicyManagedNamespaceMetadata `field:"optional" json:"managedNamespaceMetadata" yaml:"managedNamespaceMetadata"`
	Retry *ApplicationSetSpecGeneratorsMatrixGeneratorsListTemplateSpecSyncPolicyRetry `field:"optional" json:"retry" yaml:"retry"`
	SyncOptions *[]*string `field:"optional" json:"syncOptions" yaml:"syncOptions"`
}

