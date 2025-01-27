package resources

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/fordneild/omega/imports/argoprojio"
)

type ProjectProps struct {
	// Name of the project
	ProjectName string
	// Git repository URL containing the project manifests
	RepoUrl string
	// Path within the repo containing the project manifests
	Path string
	// List of child application configurations
	ChildApps []ChildAppConfig
}

type ChildAppConfig struct {
	// Name of the child application
	Name string
	// Path within the project repo for this application
	Path string
}

func NewProjectArgocdResources(scope constructs.Construct, id string, props ProjectProps) constructs.Construct {
	project := constructs.NewConstruct(scope, &id)

	// 1. Create the ArgoCD Project CRD
	argoprojio.NewAppProject(project, jsii.String(id), &argoprojio.AppProjectProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name:      jsii.String(props.ProjectName),
			Namespace: jsii.String("argocd"),
		},
		// Add any project-specific settings here
		Spec: &argoprojio.AppProjectSpec{
			ClusterResourceWhitelist: &[]*argoprojio.AppProjectSpecClusterResourceWhitelist{
				{
					Group: jsii.String("*"),
					Kind:  jsii.String("*"),
				},
			},
			Destinations: &[]*argoprojio.AppProjectSpecDestinations{
				{
					Namespace: jsii.String("*"),
					Server:    jsii.String("*"),
				},
			},
			SourceRepos: jsii.Strings("*"),
		},
	})

	// 2. Create the Project Root Application
	argoprojio.NewApplication(project, jsii.String(id+"-root"), &argoprojio.ApplicationProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Namespace: jsii.String("argocd"),
			Name:      jsii.String(props.ProjectName + "-root"),
		},
		Spec: &argoprojio.ApplicationSpec{
			Project: jsii.String(props.ProjectName),
			Source: &argoprojio.ApplicationSpecSource{
				RepoUrl: jsii.String(props.RepoUrl),
				Path:    jsii.String(props.Path),
			},
			Destination: &argoprojio.ApplicationSpecDestination{
				Server: jsii.String("https://kubernetes.default.svc"),
				// Namespace: jsii.String("*"),
			},
		},
	})

	// 3. Create Child Applications
	// for _, childApp := range props.ChildApps {
	// 	argoprojio.NewApplication(project, jsii.String(id+"-"+childApp.Name), &argoprojio.ApplicationProps{
	// 		Metadata: &cdk8s.ApiObjectMetadata{
	// 			Namespace: jsii.String("argocd"),
	// 			Name:      jsii.String(props.ProjectName + "-" + childApp.Name),
	// 		},
	// 		Spec: &argoprojio.ApplicationSpec{
	// 			Project: jsii.String(props.ProjectName),
	// 			Source: &argoprojio.ApplicationSpecSource{
	// 				RepoUrl: jsii.String(props.RepoUrl),
	// 				Path:    jsii.String(childApp.Path),
	// 			},
	// 		},
	// 	})
	// }

	return project
}
