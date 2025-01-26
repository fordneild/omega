package argoprojio


type ApplicationSetSpecSyncPolicy struct {
	ApplicationsSync ApplicationSetSpecSyncPolicyApplicationsSync `field:"optional" json:"applicationsSync" yaml:"applicationsSync"`
	PreserveResourcesOnDeletion *bool `field:"optional" json:"preserveResourcesOnDeletion" yaml:"preserveResourcesOnDeletion"`
}

