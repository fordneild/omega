package resources

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type MyChartProps struct {
	cdk8s.ChartProps
}

func NewChart(scope constructs.Construct, id string, ns string, appLabel string) cdk8s.Chart {
	chart := cdk8s.NewChart(scope, jsii.String(id), &cdk8s.ChartProps{
		Namespace: jsii.String(ns),
	})

	NewProject(chart, "project", ProjectProps{
		ProjectName: "crossplane",
		RepoUrl:     "https://github.come/fordneild/omega/projects.git",
		Path:        "crossplane",
		ChildApps: []ChildAppConfig{
			{Name: "crossplane-system", Path: "crossplane-system"},
		},
	})

	return chart
}
