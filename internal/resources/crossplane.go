package resources

import (
	"fmt"

	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/fordneild/omega/internal/project"
)

func NewAWSCrossplane(project project.Project) cdk8s.App {
	app := cdk8s.NewApp(nil)
	chart := cdk8s.NewChart(app, jsii.String(fmt.Sprintf("%s-crossplane-system-helm-chart", project.GetId())), nil)
	// crossplane helm chart
	NewHelmChartAsArgocdApp(chart, project.GetId(), ArgocdHelmChartProps{
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
	// TODO set up AWS provider
	return app
}
