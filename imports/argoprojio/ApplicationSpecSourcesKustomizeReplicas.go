package argoprojio


type ApplicationSpecSourcesKustomizeReplicas struct {
	// Number of replicas.
	Count ApplicationSpecSourcesKustomizeReplicasCount `field:"required" json:"count" yaml:"count"`
	// Name of Deployment or StatefulSet.
	Name *string `field:"required" json:"name" yaml:"name"`
}

