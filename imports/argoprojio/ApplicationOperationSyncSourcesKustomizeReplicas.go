package argoprojio


type ApplicationOperationSyncSourcesKustomizeReplicas struct {
	// Number of replicas.
	Count ApplicationOperationSyncSourcesKustomizeReplicasCount `field:"required" json:"count" yaml:"count"`
	// Name of Deployment or StatefulSet.
	Name *string `field:"required" json:"name" yaml:"name"`
}

