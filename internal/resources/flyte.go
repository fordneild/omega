package resources

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/fordneild/omega/internal/project"
)

func NewFlyteHelmChart(project project.Project) cdk8s.App {
	app := cdk8s.NewApp(nil)
	chart := cdk8s.NewChart(app, project.GetSubId("chart"), nil)
	NewHelmChartAsArgocdApp(chart, project.GetSubId("helm-chart"), ArgocdHelmChartProps{
		Name:                   jsii.String("flyte-helm-chart"),
		ArgocdNamespace:        jsii.String("argocd"),
		ReleaseNamespace:       jsii.String("flyte"),
		ProjectName:            jsii.String(project.GetId()),
		ChartName:              jsii.String("flyte"),
		ChartRepoUrl:           jsii.String("https://flyteorg.github.io/flyte"),
		ChartTargetRevision:    jsii.String("0.1.10"),
		ReleaseName:            jsii.String("flyte"),
		DisableAutomatedSync:   jsii.Bool(false),
		DisableCreateNamespace: jsii.Bool(false),
		PruneOnAutomatedSync:   jsii.Bool(true),
	})
	return app

}
