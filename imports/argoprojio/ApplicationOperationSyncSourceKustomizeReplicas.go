package argoprojio


type ApplicationOperationSyncSourceKustomizeReplicas struct {
	// Number of replicas.
	Count ApplicationOperationSyncSourceKustomizeReplicasCount `field:"required" json:"count" yaml:"count"`
	// Name of Deployment or StatefulSet.
	Name *string `field:"required" json:"name" yaml:"name"`
}

