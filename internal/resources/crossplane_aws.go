package resources

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	awsupboundio "github.com/fordneild/omega/imports/awsupboundio"
	crossplane "github.com/fordneild/omega/imports/pkgcrossplaneio"
	"github.com/fordneild/omega/internal/project"
)

const (
	// only used for EKS provider for now
	CrossplaneAwsProviderVersion = "v1.21.1"
	CrossplaneHelmChartVersion   = "1.19.1"
)

// requires a secret to be created like so: https://docs.crossplane.io/v1.19/getting-started/provider-aws/#create-a-kubernetes-secret-with-the-aws-credentials
func NewAWSCrossplane(project project.Project) cdk8s.App {
	app := cdk8s.NewApp(nil)
	crossplaneAwsChart := cdk8s.NewChart(app, project.GetSubId("chart"), nil)
	// crossplane helm chart
	NewHelmChartAsArgocdApp(crossplaneAwsChart, project.GetSubId("crossplane-system-helm-chart"), ArgocdHelmChartProps{
		Name:                   jsii.String("crossplane-system-helm-chart"),
		ArgocdNamespace:        jsii.String("argocd"),
		ReleaseNamespace:       jsii.String("crossplane-system"),
		ProjectName:            jsii.String(project.GetId()),
		ChartName:              jsii.String("crossplane"),
		ChartRepoUrl:           jsii.String("https://charts.crossplane.io/stable"),
		ChartTargetRevision:    jsii.String(CrossplaneHelmChartVersion),
		ReleaseName:            jsii.String("crossplane-system"),
		DisableAutomatedSync:   jsii.Bool(false),
		DisableCreateNamespace: jsii.Bool(false),
		PruneOnAutomatedSync:   jsii.Bool(true),
	})
	// Set up AWS EKS provider
	crossplane.NewProvider(crossplaneAwsChart, project.GetSubId("provider-aws"), &crossplane.ProviderProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name: jsii.String("provider-aws"),
		},
		Spec: &crossplane.ProviderSpec{
			Package: jsii.Sprintf("xpkg.crossplane.io/crossplane-contrib/provider-aws-eks:%s", CrossplaneAwsProviderVersion),
		},
	})
	// Set up AWS EKS provider config
	awsupboundio.NewProviderConfig(crossplaneAwsChart, project.GetSubId("provider-config-aws"), &awsupboundio.ProviderConfigProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name: jsii.String("provider-config-aws"),
		},
		Spec: &awsupboundio.ProviderConfigSpec{
			Credentials: &awsupboundio.ProviderConfigSpecCredentials{
				Source: awsupboundio.ProviderConfigSpecCredentialsSource_SECRET,
				SecretRef: &awsupboundio.ProviderConfigSpecCredentialsSecretRef{
					Namespace: jsii.String("crossplane-system"),
					Name:      jsii.String("aws-secret"),
					Key:       jsii.String("creds"),
				},
			},
		},
	})
	return app
}
