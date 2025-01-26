package argoprojio


// SyncPolicy controls when and how a sync will be performed.
type ApplicationSpecSyncPolicy struct {
	// Automated will keep an application synced to the target revision.
	Automated *ApplicationSpecSyncPolicyAutomated `field:"optional" json:"automated" yaml:"automated"`
	// ManagedNamespaceMetadata controls metadata in the given namespace (if CreateNamespace=true).
	ManagedNamespaceMetadata *ApplicationSpecSyncPolicyManagedNamespaceMetadata `field:"optional" json:"managedNamespaceMetadata" yaml:"managedNamespaceMetadata"`
	// Retry controls failed sync retry behavior.
	Retry *ApplicationSpecSyncPolicyRetry `field:"optional" json:"retry" yaml:"retry"`
	// Options allow you to specify whole app sync-options.
	SyncOptions *[]*string `field:"optional" json:"syncOptions" yaml:"syncOptions"`
}

