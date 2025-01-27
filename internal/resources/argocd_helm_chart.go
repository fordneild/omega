package resources

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/fordneild/omega/imports/argoprojio"
	"github.com/samber/lo"
)

type ArgocdHelmChartProps struct {
	Name                *string
	ArgocdNamespace     *string
	ReleaseNamespace    *string
	ProjectName         *string
	ChartName           *string
	ChartRepoUrl        *string
	ChartTargetRevision *string
	ReleaseName         *string
}

func NewArgocdHelmChart(scope constructs.Construct, id string, props ArgocdHelmChartProps) constructs.Construct {
	project := constructs.NewConstruct(scope, &id)
	// 2. Create the Project Root Application
	argocdNamespace, _ := lo.Coalesce(props.ArgocdNamespace, jsii.String("argocd"))
	argoprojio.NewApplication(project, jsii.String(id), &argoprojio.ApplicationProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name:      props.Name,
			Namespace: argocdNamespace,
		},
		Spec: &argoprojio.ApplicationSpec{
			Project: jsii.String(*props.ProjectName),
			Source: &argoprojio.ApplicationSpecSource{
				Chart:          props.ChartName,
				RepoUrl:        props.ChartRepoUrl,
				TargetRevision: props.ChartTargetRevision,
				Helm: &argoprojio.ApplicationSpecSourceHelm{
					ReleaseName: props.ReleaseName,
					// ValueFiles: jsii.Strings("values.yaml"),
				},
			},
			Destination: &argoprojio.ApplicationSpecDestination{
				Server:    jsii.String("https://kubernetes.default.svc"),
				Namespace: props.ReleaseNamespace,
			},
		},
	})

	return project
}
