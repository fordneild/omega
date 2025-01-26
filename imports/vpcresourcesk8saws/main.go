// vpcresourcesk8saws
package vpcresourcesk8saws

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"vpcresourcesk8saws.CniNode",
		reflect.TypeOf((*CniNode)(nil)).Elem(),
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
			j := jsiiProxy_CniNode{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"vpcresourcesk8saws.CniNodeProps",
		reflect.TypeOf((*CniNodeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"vpcresourcesk8saws.CniNodeSpec",
		reflect.TypeOf((*CniNodeSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"vpcresourcesk8saws.CniNodeSpecFeatures",
		reflect.TypeOf((*CniNodeSpecFeatures)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"vpcresourcesk8saws.SecurityGroupPolicy",
		reflect.TypeOf((*SecurityGroupPolicy)(nil)).Elem(),
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
			j := jsiiProxy_SecurityGroupPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"vpcresourcesk8saws.SecurityGroupPolicyProps",
		reflect.TypeOf((*SecurityGroupPolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"vpcresourcesk8saws.SecurityGroupPolicySpec",
		reflect.TypeOf((*SecurityGroupPolicySpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"vpcresourcesk8saws.SecurityGroupPolicySpecPodSelector",
		reflect.TypeOf((*SecurityGroupPolicySpecPodSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"vpcresourcesk8saws.SecurityGroupPolicySpecPodSelectorMatchExpressions",
		reflect.TypeOf((*SecurityGroupPolicySpecPodSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"vpcresourcesk8saws.SecurityGroupPolicySpecSecurityGroups",
		reflect.TypeOf((*SecurityGroupPolicySpecSecurityGroups)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"vpcresourcesk8saws.SecurityGroupPolicySpecServiceAccountSelector",
		reflect.TypeOf((*SecurityGroupPolicySpecServiceAccountSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"vpcresourcesk8saws.SecurityGroupPolicySpecServiceAccountSelectorMatchExpressions",
		reflect.TypeOf((*SecurityGroupPolicySpecServiceAccountSelectorMatchExpressions)(nil)).Elem(),
	)
}
