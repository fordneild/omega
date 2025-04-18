package resources

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	crossplane "github.com/fordneild/omega/imports/pkgcrossplaneio"
	"github.com/fordneild/omega/internal/project"
)

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
		ChartTargetRevision:    jsii.String("1.10.0"),
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
			Package: jsii.String("xpkg.crossplane.io/crossplane-contrib/provider-aws-eks:v1.21.1"),
		},
	})
	return app
}
