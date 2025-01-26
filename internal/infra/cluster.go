package infra

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const argocdNamespace = "argocd"

type Cluster struct {
}

var _ PulumiHook = &Cluster{}

func NewCluster() *Cluster {
	return &Cluster{}
}

func (h *Cluster) Run(ctx *pulumi.Context) error {
	eks, err := NewEKSCluster().Run(ctx)
	if err != nil {
		return err
	}
	eksProvider, err := kubernetes.NewProvider(ctx, "eks-provider", &kubernetes.ProviderArgs{
		Kubeconfig: eks.KubeconfigJson,
	})
	if err != nil {
		return err
	}
	argocdComponent, err := NewArgocd(ctx, "argocd", ArgocdComponentArgs{
		namespace: argocdNamespace,
		provider:  eksProvider,
	}, pulumi.Provider(eksProvider))
	if err != nil {
		return err
	}
	_, err = NewRootApp(ctx, RootAppArgs{
		eksProvider:     eksProvider,
		argocdComponent: argocdComponent,
		namespace:       argocdNamespace,
	})
	if err != nil {
		return err
	}
	return nil
}
