package argoprojio


type ApplicationSpecSourceKustomizeReplicas struct {
	// Number of replicas.
	Count ApplicationSpecSourceKustomizeReplicasCount `field:"required" json:"count" yaml:"count"`
	// Name of Deployment or StatefulSet.
	Name *string `field:"required" json:"name" yaml:"name"`
}

