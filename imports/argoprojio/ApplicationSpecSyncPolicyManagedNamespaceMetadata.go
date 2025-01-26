package argoprojio


// ManagedNamespaceMetadata controls metadata in the given namespace (if CreateNamespace=true).
type ApplicationSpecSyncPolicyManagedNamespaceMetadata struct {
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

