// argoprojio
package argoprojio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"argoprojio.AppProject",
		reflect.TypeOf((*AppProject)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AppProject{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"argoprojio.AppProjectProps",
		reflect.TypeOf((*AppProjectProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.AppProjectSpec",
		reflect.TypeOf((*AppProjectSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.AppProjectSpecClusterResourceBlacklist",
		reflect.TypeOf((*AppProjectSpecClusterResourceBlacklist)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.AppProjectSpecClusterResourceWhitelist",
		reflect.TypeOf((*AppProjectSpecClusterResourceWhitelist)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.AppProjectSpecDestinationServiceAccounts",
		reflect.TypeOf((*AppProjectSpecDestinationServiceAccounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.AppProjectSpecDestinations",
		reflect.TypeOf((*AppProjectSpecDestinations)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.AppProjectSpecNamespaceResourceBlacklist",
		reflect.TypeOf((*AppProjectSpecNamespaceResourceBlacklist)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.AppProjectSpecNamespaceResourceWhitelist",
		reflect.TypeOf((*AppProjectSpecNamespaceResourceWhitelist)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.AppProjectSpecOrphanedResources",
		reflect.TypeOf((*AppProjectSpecOrphanedResources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.AppProjectSpecOrphanedResourcesIgnore",
		reflect.TypeOf((*AppProjectSpecOrphanedResourcesIgnore)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.AppProjectSpecRoles",
		reflect.TypeOf((*AppProjectSpecRoles)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.AppProjectSpecRolesJwtTokens",
		reflect.TypeOf((*AppProjectSpecRolesJwtTokens)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.AppProjectSpecSignatureKeys",
		reflect.TypeOf((*AppProjectSpecSignatureKeys)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.AppProjectSpecSyncWindows",
		reflect.TypeOf((*AppProjectSpecSyncWindows)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"argoprojio.Application",
		reflect.TypeOf((*Application)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Application{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperation",
		reflect.TypeOf((*ApplicationOperation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationInfo",
		reflect.TypeOf((*ApplicationOperationInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationInitiatedBy",
		reflect.TypeOf((*ApplicationOperationInitiatedBy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationRetry",
		reflect.TypeOf((*ApplicationOperationRetry)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationRetryBackoff",
		reflect.TypeOf((*ApplicationOperationRetryBackoff)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSync",
		reflect.TypeOf((*ApplicationOperationSync)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncResources",
		reflect.TypeOf((*ApplicationOperationSyncResources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSource",
		reflect.TypeOf((*ApplicationOperationSyncSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSourceDirectory",
		reflect.TypeOf((*ApplicationOperationSyncSourceDirectory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSourceDirectoryJsonnet",
		reflect.TypeOf((*ApplicationOperationSyncSourceDirectoryJsonnet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSourceDirectoryJsonnetExtVars",
		reflect.TypeOf((*ApplicationOperationSyncSourceDirectoryJsonnetExtVars)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSourceDirectoryJsonnetTlas",
		reflect.TypeOf((*ApplicationOperationSyncSourceDirectoryJsonnetTlas)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSourceHelm",
		reflect.TypeOf((*ApplicationOperationSyncSourceHelm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSourceHelmFileParameters",
		reflect.TypeOf((*ApplicationOperationSyncSourceHelmFileParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSourceHelmParameters",
		reflect.TypeOf((*ApplicationOperationSyncSourceHelmParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSourceKustomize",
		reflect.TypeOf((*ApplicationOperationSyncSourceKustomize)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSourceKustomizePatches",
		reflect.TypeOf((*ApplicationOperationSyncSourceKustomizePatches)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSourceKustomizePatchesTarget",
		reflect.TypeOf((*ApplicationOperationSyncSourceKustomizePatchesTarget)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSourceKustomizeReplicas",
		reflect.TypeOf((*ApplicationOperationSyncSourceKustomizeReplicas)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"argoprojio.ApplicationOperationSyncSourceKustomizeReplicasCount",
		reflect.TypeOf((*ApplicationOperationSyncSourceKustomizeReplicasCount)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ApplicationOperationSyncSourceKustomizeReplicasCount{}
		},
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSourcePlugin",
		reflect.TypeOf((*ApplicationOperationSyncSourcePlugin)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSourcePluginEnv",
		reflect.TypeOf((*ApplicationOperationSyncSourcePluginEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSourcePluginParameters",
		reflect.TypeOf((*ApplicationOperationSyncSourcePluginParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSources",
		reflect.TypeOf((*ApplicationOperationSyncSources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSourcesDirectory",
		reflect.TypeOf((*ApplicationOperationSyncSourcesDirectory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSourcesDirectoryJsonnet",
		reflect.TypeOf((*ApplicationOperationSyncSourcesDirectoryJsonnet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSourcesDirectoryJsonnetExtVars",
		reflect.TypeOf((*ApplicationOperationSyncSourcesDirectoryJsonnetExtVars)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSourcesDirectoryJsonnetTlas",
		reflect.TypeOf((*ApplicationOperationSyncSourcesDirectoryJsonnetTlas)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSourcesHelm",
		reflect.TypeOf((*ApplicationOperationSyncSourcesHelm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSourcesHelmFileParameters",
		reflect.TypeOf((*ApplicationOperationSyncSourcesHelmFileParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSourcesHelmParameters",
		reflect.TypeOf((*ApplicationOperationSyncSourcesHelmParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSourcesKustomize",
		reflect.TypeOf((*ApplicationOperationSyncSourcesKustomize)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSourcesKustomizePatches",
		reflect.TypeOf((*ApplicationOperationSyncSourcesKustomizePatches)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSourcesKustomizePatchesTarget",
		reflect.TypeOf((*ApplicationOperationSyncSourcesKustomizePatchesTarget)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSourcesKustomizeReplicas",
		reflect.TypeOf((*ApplicationOperationSyncSourcesKustomizeReplicas)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"argoprojio.ApplicationOperationSyncSourcesKustomizeReplicasCount",
		reflect.TypeOf((*ApplicationOperationSyncSourcesKustomizeReplicasCount)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ApplicationOperationSyncSourcesKustomizeReplicasCount{}
		},
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSourcesPlugin",
		reflect.TypeOf((*ApplicationOperationSyncSourcesPlugin)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSourcesPluginEnv",
		reflect.TypeOf((*ApplicationOperationSyncSourcesPluginEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSourcesPluginParameters",
		reflect.TypeOf((*ApplicationOperationSyncSourcesPluginParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSyncStrategy",
		reflect.TypeOf((*ApplicationOperationSyncSyncStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSyncStrategyApply",
		reflect.TypeOf((*ApplicationOperationSyncSyncStrategyApply)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationOperationSyncSyncStrategyHook",
		reflect.TypeOf((*ApplicationOperationSyncSyncStrategyHook)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationProps",
		reflect.TypeOf((*ApplicationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpec",
		reflect.TypeOf((*ApplicationSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecDestination",
		reflect.TypeOf((*ApplicationSpecDestination)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecIgnoreDifferences",
		reflect.TypeOf((*ApplicationSpecIgnoreDifferences)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecInfo",
		reflect.TypeOf((*ApplicationSpecInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSource",
		reflect.TypeOf((*ApplicationSpecSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSourceDirectory",
		reflect.TypeOf((*ApplicationSpecSourceDirectory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSourceDirectoryJsonnet",
		reflect.TypeOf((*ApplicationSpecSourceDirectoryJsonnet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSourceDirectoryJsonnetExtVars",
		reflect.TypeOf((*ApplicationSpecSourceDirectoryJsonnetExtVars)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSourceDirectoryJsonnetTlas",
		reflect.TypeOf((*ApplicationSpecSourceDirectoryJsonnetTlas)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSourceHelm",
		reflect.TypeOf((*ApplicationSpecSourceHelm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSourceHelmFileParameters",
		reflect.TypeOf((*ApplicationSpecSourceHelmFileParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSourceHelmParameters",
		reflect.TypeOf((*ApplicationSpecSourceHelmParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSourceKustomize",
		reflect.TypeOf((*ApplicationSpecSourceKustomize)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSourceKustomizePatches",
		reflect.TypeOf((*ApplicationSpecSourceKustomizePatches)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSourceKustomizePatchesTarget",
		reflect.TypeOf((*ApplicationSpecSourceKustomizePatchesTarget)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSourceKustomizeReplicas",
		reflect.TypeOf((*ApplicationSpecSourceKustomizeReplicas)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"argoprojio.ApplicationSpecSourceKustomizeReplicasCount",
		reflect.TypeOf((*ApplicationSpecSourceKustomizeReplicasCount)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ApplicationSpecSourceKustomizeReplicasCount{}
		},
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSourcePlugin",
		reflect.TypeOf((*ApplicationSpecSourcePlugin)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSourcePluginEnv",
		reflect.TypeOf((*ApplicationSpecSourcePluginEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSourcePluginParameters",
		reflect.TypeOf((*ApplicationSpecSourcePluginParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSources",
		reflect.TypeOf((*ApplicationSpecSources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSourcesDirectory",
		reflect.TypeOf((*ApplicationSpecSourcesDirectory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSourcesDirectoryJsonnet",
		reflect.TypeOf((*ApplicationSpecSourcesDirectoryJsonnet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSourcesDirectoryJsonnetExtVars",
		reflect.TypeOf((*ApplicationSpecSourcesDirectoryJsonnetExtVars)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSourcesDirectoryJsonnetTlas",
		reflect.TypeOf((*ApplicationSpecSourcesDirectoryJsonnetTlas)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSourcesHelm",
		reflect.TypeOf((*ApplicationSpecSourcesHelm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSourcesHelmFileParameters",
		reflect.TypeOf((*ApplicationSpecSourcesHelmFileParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSourcesHelmParameters",
		reflect.TypeOf((*ApplicationSpecSourcesHelmParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSourcesKustomize",
		reflect.TypeOf((*ApplicationSpecSourcesKustomize)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSourcesKustomizePatches",
		reflect.TypeOf((*ApplicationSpecSourcesKustomizePatches)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSourcesKustomizePatchesTarget",
		reflect.TypeOf((*ApplicationSpecSourcesKustomizePatchesTarget)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSourcesKustomizeReplicas",
		reflect.TypeOf((*ApplicationSpecSourcesKustomizeReplicas)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"argoprojio.ApplicationSpecSourcesKustomizeReplicasCount",
		reflect.TypeOf((*ApplicationSpecSourcesKustomizeReplicasCount)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ApplicationSpecSourcesKustomizeReplicasCount{}
		},
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSourcesPlugin",
		reflect.TypeOf((*ApplicationSpecSourcesPlugin)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSourcesPluginEnv",
		reflect.TypeOf((*ApplicationSpecSourcesPluginEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSourcesPluginParameters",
		reflect.TypeOf((*ApplicationSpecSourcesPluginParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSyncPolicy",
		reflect.TypeOf((*ApplicationSpecSyncPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSyncPolicyAutomated",
		reflect.TypeOf((*ApplicationSpecSyncPolicyAutomated)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSyncPolicyManagedNamespaceMetadata",
		reflect.TypeOf((*ApplicationSpecSyncPolicyManagedNamespaceMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSyncPolicyRetry",
		reflect.TypeOf((*ApplicationSpecSyncPolicyRetry)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"argoprojio.ApplicationSpecSyncPolicyRetryBackoff",
		reflect.TypeOf((*ApplicationSpecSyncPolicyRetryBackoff)(nil)).Elem(),
	)
}
