package argoprojio


// ApplicationSource contains all required information about the source of an application.
type ApplicationSpecSources struct {
	// RepoURL is the URL to the repository (Git or Helm) that contains the application manifests.
	RepoUrl *string `field:"required" json:"repoUrl" yaml:"repoUrl"`
	// Chart is a Helm chart name, and must be specified for applications sourced from a Helm repo.
	Chart *string `field:"optional" json:"chart" yaml:"chart"`
	// Directory holds path/directory specific options.
	Directory *ApplicationSpecSourcesDirectory `field:"optional" json:"directory" yaml:"directory"`
	// Helm holds helm specific options.
	Helm *ApplicationSpecSourcesHelm `field:"optional" json:"helm" yaml:"helm"`
	// Kustomize holds kustomize specific options.
	Kustomize *ApplicationSpecSourcesKustomize `field:"optional" json:"kustomize" yaml:"kustomize"`
	// Path is a directory path within the Git repository, and is only valid for applications sourced from Git.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Plugin holds config management plugin specific options.
	Plugin *ApplicationSpecSourcesPlugin `field:"optional" json:"plugin" yaml:"plugin"`
	// Ref is reference to another source within sources field.
	//
	// This field will not be used if used with a `source` tag.
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
	// TargetRevision defines the revision of the source to sync the application to.
	//
	// In case of Git, this can be commit, tag, or branch. If omitted, will equal to HEAD.
	// In case of Helm, this is a semver tag for the Chart's version.
	TargetRevision *string `field:"optional" json:"targetRevision" yaml:"targetRevision"`
}

