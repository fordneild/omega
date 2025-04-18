// awsupboundio
package awsupboundio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"awsupboundio.ProviderConfig",
		reflect.TypeOf((*ProviderConfig)(nil)).Elem(),
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
			j := jsiiProxy_ProviderConfig{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awsupboundio.ProviderConfigProps",
		reflect.TypeOf((*ProviderConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awsupboundio.ProviderConfigSpec",
		reflect.TypeOf((*ProviderConfigSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awsupboundio.ProviderConfigSpecAssumeRoleChain",
		reflect.TypeOf((*ProviderConfigSpecAssumeRoleChain)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awsupboundio.ProviderConfigSpecAssumeRoleChainTags",
		reflect.TypeOf((*ProviderConfigSpecAssumeRoleChainTags)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awsupboundio.ProviderConfigSpecCredentials",
		reflect.TypeOf((*ProviderConfigSpecCredentials)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awsupboundio.ProviderConfigSpecCredentialsEnv",
		reflect.TypeOf((*ProviderConfigSpecCredentialsEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awsupboundio.ProviderConfigSpecCredentialsFs",
		reflect.TypeOf((*ProviderConfigSpecCredentialsFs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awsupboundio.ProviderConfigSpecCredentialsSecretRef",
		reflect.TypeOf((*ProviderConfigSpecCredentialsSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awsupboundio.ProviderConfigSpecCredentialsSource",
		reflect.TypeOf((*ProviderConfigSpecCredentialsSource)(nil)).Elem(),
		map[string]interface{}{
			"NONE": ProviderConfigSpecCredentialsSource_NONE,
			"SECRET": ProviderConfigSpecCredentialsSource_SECRET,
			"IRSA": ProviderConfigSpecCredentialsSource_IRSA,
			"WEB_IDENTITY": ProviderConfigSpecCredentialsSource_WEB_IDENTITY,
			"POD_IDENTITY": ProviderConfigSpecCredentialsSource_POD_IDENTITY,
			"UPBOUND": ProviderConfigSpecCredentialsSource_UPBOUND,
		},
	)
	_jsii_.RegisterStruct(
		"awsupboundio.ProviderConfigSpecCredentialsUpbound",
		reflect.TypeOf((*ProviderConfigSpecCredentialsUpbound)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awsupboundio.ProviderConfigSpecCredentialsUpboundWebIdentity",
		reflect.TypeOf((*ProviderConfigSpecCredentialsUpboundWebIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awsupboundio.ProviderConfigSpecCredentialsUpboundWebIdentityTokenConfig",
		reflect.TypeOf((*ProviderConfigSpecCredentialsUpboundWebIdentityTokenConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awsupboundio.ProviderConfigSpecCredentialsUpboundWebIdentityTokenConfigFs",
		reflect.TypeOf((*ProviderConfigSpecCredentialsUpboundWebIdentityTokenConfigFs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awsupboundio.ProviderConfigSpecCredentialsUpboundWebIdentityTokenConfigSecretRef",
		reflect.TypeOf((*ProviderConfigSpecCredentialsUpboundWebIdentityTokenConfigSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awsupboundio.ProviderConfigSpecCredentialsUpboundWebIdentityTokenConfigSource",
		reflect.TypeOf((*ProviderConfigSpecCredentialsUpboundWebIdentityTokenConfigSource)(nil)).Elem(),
		map[string]interface{}{
			"SECRET": ProviderConfigSpecCredentialsUpboundWebIdentityTokenConfigSource_SECRET,
			"FILESYSTEM": ProviderConfigSpecCredentialsUpboundWebIdentityTokenConfigSource_FILESYSTEM,
		},
	)
	_jsii_.RegisterStruct(
		"awsupboundio.ProviderConfigSpecCredentialsWebIdentity",
		reflect.TypeOf((*ProviderConfigSpecCredentialsWebIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awsupboundio.ProviderConfigSpecCredentialsWebIdentityTokenConfig",
		reflect.TypeOf((*ProviderConfigSpecCredentialsWebIdentityTokenConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awsupboundio.ProviderConfigSpecCredentialsWebIdentityTokenConfigFs",
		reflect.TypeOf((*ProviderConfigSpecCredentialsWebIdentityTokenConfigFs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awsupboundio.ProviderConfigSpecCredentialsWebIdentityTokenConfigSecretRef",
		reflect.TypeOf((*ProviderConfigSpecCredentialsWebIdentityTokenConfigSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awsupboundio.ProviderConfigSpecCredentialsWebIdentityTokenConfigSource",
		reflect.TypeOf((*ProviderConfigSpecCredentialsWebIdentityTokenConfigSource)(nil)).Elem(),
		map[string]interface{}{
			"SECRET": ProviderConfigSpecCredentialsWebIdentityTokenConfigSource_SECRET,
			"FILESYSTEM": ProviderConfigSpecCredentialsWebIdentityTokenConfigSource_FILESYSTEM,
		},
	)
	_jsii_.RegisterStruct(
		"awsupboundio.ProviderConfigSpecEndpoint",
		reflect.TypeOf((*ProviderConfigSpecEndpoint)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awsupboundio.ProviderConfigSpecEndpointSource",
		reflect.TypeOf((*ProviderConfigSpecEndpointSource)(nil)).Elem(),
		map[string]interface{}{
			"SERVICE_METADATA": ProviderConfigSpecEndpointSource_SERVICE_METADATA,
			"CUSTOM": ProviderConfigSpecEndpointSource_CUSTOM,
		},
	)
	_jsii_.RegisterStruct(
		"awsupboundio.ProviderConfigSpecEndpointUrl",
		reflect.TypeOf((*ProviderConfigSpecEndpointUrl)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awsupboundio.ProviderConfigSpecEndpointUrlDynamic",
		reflect.TypeOf((*ProviderConfigSpecEndpointUrlDynamic)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awsupboundio.ProviderConfigSpecEndpointUrlDynamicProtocol",
		reflect.TypeOf((*ProviderConfigSpecEndpointUrlDynamicProtocol)(nil)).Elem(),
		map[string]interface{}{
			"HTTP": ProviderConfigSpecEndpointUrlDynamicProtocol_HTTP,
			"HTTPS": ProviderConfigSpecEndpointUrlDynamicProtocol_HTTPS,
		},
	)
	_jsii_.RegisterEnum(
		"awsupboundio.ProviderConfigSpecEndpointUrlType",
		reflect.TypeOf((*ProviderConfigSpecEndpointUrlType)(nil)).Elem(),
		map[string]interface{}{
			"STATIC": ProviderConfigSpecEndpointUrlType_STATIC,
			"DYNAMIC": ProviderConfigSpecEndpointUrlType_DYNAMIC,
			"AUTO": ProviderConfigSpecEndpointUrlType_AUTO,
		},
	)
}
