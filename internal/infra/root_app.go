package infra

import (
	"fmt"

	"github.com/fordneild/omega/internal/project"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/apiextensions"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RootAppArgs struct {
	eksProvider     *kubernetes.Provider
	argocdComponent *ArgocdComponent
	namespace       string
}

func NewRootApp(ctx *pulumi.Context, args RootAppArgs) (*apiextensions.CustomResource, error) {
	return apiextensions.NewCustomResource(ctx, "plain-cr", &apiextensions.CustomResourceArgs{
		ApiVersion: pulumi.String("argoproj.io/v1alpha1"),
		Kind:       pulumi.String("Application"),
		Metadata: &metav1.ObjectMetaArgs{
			Namespace:   pulumi.String(args.namespace),
			Name:        pulumi.String("root-app-of-apps"),
			Annotations: pulumi.StringMap{},
		},
		OtherFields: kubernetes.UntypedArgs{
			"spec": kubernetes.UntypedArgs{
				"project": "default",
				"source": kubernetes.UntypedArgs{
					"repoURL":        "https://github.com/fordneild/omega.git",
					"targetRevision": "HEAD",
					"path":           fmt.Sprintf("%s/root", project.OmegaProjectBasePath),
				},
				"destination": kubernetes.UntypedArgs{
					"server":    "https://kubernetes.default.svc",
					"namespace": args.namespace,
				},
				"syncPolicy": kubernetes.UntypedArgs{
					"automated": kubernetes.UntypedArgs{
						"prune": true,
					},
					"syncOptions": []string{"CreateNamespace=true"},
				},
			},
		},
	}, pulumi.Provider(args.eksProvider), pulumi.DependsOn([]pulumi.Resource{args.argocdComponent}))
}
