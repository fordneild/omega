// networkingk8saws
package networkingk8saws

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"networkingk8saws.PolicyEndpoint",
		reflect.TypeOf((*PolicyEndpoint)(nil)).Elem(),
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
			j := jsiiProxy_PolicyEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"networkingk8saws.PolicyEndpointProps",
		reflect.TypeOf((*PolicyEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingk8saws.PolicyEndpointSpec",
		reflect.TypeOf((*PolicyEndpointSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingk8saws.PolicyEndpointSpecEgress",
		reflect.TypeOf((*PolicyEndpointSpecEgress)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingk8saws.PolicyEndpointSpecEgressPorts",
		reflect.TypeOf((*PolicyEndpointSpecEgressPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingk8saws.PolicyEndpointSpecIngress",
		reflect.TypeOf((*PolicyEndpointSpecIngress)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingk8saws.PolicyEndpointSpecIngressPorts",
		reflect.TypeOf((*PolicyEndpointSpecIngressPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingk8saws.PolicyEndpointSpecPodSelector",
		reflect.TypeOf((*PolicyEndpointSpecPodSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingk8saws.PolicyEndpointSpecPodSelectorEndpoints",
		reflect.TypeOf((*PolicyEndpointSpecPodSelectorEndpoints)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingk8saws.PolicyEndpointSpecPodSelectorMatchExpressions",
		reflect.TypeOf((*PolicyEndpointSpecPodSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"networkingk8saws.PolicyEndpointSpecPolicyRef",
		reflect.TypeOf((*PolicyEndpointSpecPolicyRef)(nil)).Elem(),
	)
}
