package infra

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	v1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"

	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

//	type Argocd struct {
//		k8sProvider kubernetes.Provider
//		namespace   string
//	}
type ArgocdComponent struct {
	pulumi.ResourceState
}
type ArgocdComponentArgs struct {
	namespace string
	provider  *kubernetes.Provider
}

func NewArgocd(ctx *pulumi.Context, name string, argocdArgs ArgocdComponentArgs, opts ...pulumi.ResourceOption) (*ArgocdComponent, error) {
	argocdComponent := ArgocdComponent{}
	namespace := argocdArgs.namespace
	provider := pulumi.Provider(argocdArgs.provider)
	parent := pulumi.Parent(&argocdComponent)
	if namespace == "" {
		namespace = "argocd"
	}
	namespaceResource, err := v1.NewNamespace(ctx, namespace, &v1.NamespaceArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Name: pulumi.String(namespace),
		},
	}, parent, provider)
	if err != nil {
		return nil, err
	}
	// cmData, err := json.Marshal()
	// if err != nil {
	// 	return nil, err
	// }
	_, err = helm.NewChart(ctx, "argocd", helm.ChartArgs{
		Chart:     pulumi.String("argo-cd"),
		Version:   pulumi.String("7.7.17"),
		Namespace: pulumi.String(namespace),
		FetchArgs: helm.FetchArgs{
			Repo: pulumi.String("https://argoproj.github.io/argo-helm"),
		},
		Values: pulumi.Map{
			"configs": pulumi.Map{
				"cm": pulumi.Map{
					// this is also to make crossplane happy. See https://docs.crossplane.io/latest/guides/crossplane-with-argo-cd/
					"application.resourceTrackingMethod": pulumi.String("annotation"),
					"resource.customizations":            pulumi.String(crossplaneArgocdResourceCustomizations),
					"resource.exclusions":                pulumi.String(crossplaneArgocdResoureExclusions),
				},
			},
		},
	}, parent, provider, pulumi.DependsOn([]pulumi.Resource{namespaceResource}))
	if err != nil {
		return nil, err
	}
	err = ctx.RegisterComponentResource("pkg:index:ArgocdComponent", name, &argocdComponent, opts...)
	if err != nil {
		return nil, err
	}
	return &argocdComponent, nil

}
