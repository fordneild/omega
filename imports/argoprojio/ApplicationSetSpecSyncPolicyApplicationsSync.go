package argoprojio


type ApplicationSetSpecSyncPolicyApplicationsSync string

const (
	// create-only.
	ApplicationSetSpecSyncPolicyApplicationsSync_CREATE_HYPHEN_ONLY ApplicationSetSpecSyncPolicyApplicationsSync = "CREATE_HYPHEN_ONLY"
	// create-update.
	ApplicationSetSpecSyncPolicyApplicationsSync_CREATE_HYPHEN_UPDATE ApplicationSetSpecSyncPolicyApplicationsSync = "CREATE_HYPHEN_UPDATE"
	// create-delete.
	ApplicationSetSpecSyncPolicyApplicationsSync_CREATE_HYPHEN_DELETE ApplicationSetSpecSyncPolicyApplicationsSync = "CREATE_HYPHEN_DELETE"
	// sync.
	ApplicationSetSpecSyncPolicyApplicationsSync_SYNC ApplicationSetSpecSyncPolicyApplicationsSync = "SYNC"
)

