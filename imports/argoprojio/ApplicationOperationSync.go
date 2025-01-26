package argoprojio


// Sync contains parameters for the operation.
type ApplicationOperationSync struct {
	// SelfHealAttemptsCount contains the number of auto-heal attempts.
	AutoHealAttemptsCount *float64 `field:"optional" json:"autoHealAttemptsCount" yaml:"autoHealAttemptsCount"`
	// DryRun specifies to perform a `kubectl apply --dry-run` without actually performing the sync.
	DryRun *bool `field:"optional" json:"dryRun" yaml:"dryRun"`
	// Manifests is an optional field that overrides sync source with a local directory for development.
	Manifests *[]*string `field:"optional" json:"manifests" yaml:"manifests"`
	// Prune specifies to delete resources from the cluster that are no longer tracked in git.
	Prune *bool `field:"optional" json:"prune" yaml:"prune"`
	// Resources describes which resources shall be part of the sync.
	Resources *[]*ApplicationOperationSyncResources `field:"optional" json:"resources" yaml:"resources"`
	// Revision is the revision (Git) or chart version (Helm) which to sync the application to If omitted, will use the revision specified in app spec.
	Revision *string `field:"optional" json:"revision" yaml:"revision"`
	// Revisions is the list of revision (Git) or chart version (Helm) which to sync each source in sources field for the application to If omitted, will use the revision specified in app spec.
	Revisions *[]*string `field:"optional" json:"revisions" yaml:"revisions"`
	// Source overrides the source definition set in the application.
	//
	// This is typically set in a Rollback operation and is nil during a Sync operation.
	Source *ApplicationOperationSyncSource `field:"optional" json:"source" yaml:"source"`
	// Sources overrides the source definition set in the application.
	//
	// This is typically set in a Rollback operation and is nil during a Sync operation.
	Sources *[]*ApplicationOperationSyncSources `field:"optional" json:"sources" yaml:"sources"`
	// SyncOptions provide per-sync sync-options, e.g. Validate=false.
	SyncOptions *[]*string `field:"optional" json:"syncOptions" yaml:"syncOptions"`
	// SyncStrategy describes how to perform the sync.
	SyncStrategy *ApplicationOperationSyncSyncStrategy `field:"optional" json:"syncStrategy" yaml:"syncStrategy"`
}

