package commands

type AllCommands struct {
	Cdk8sCommands
	PulumiCommands
	InstallCommands
}

type Cdk8sCommands struct {
	Import ImportResourceDefinitionsCmd `cmd:"" name:"import" help:"Imports K8s resource definitions from a cluster/url"`
	Render RenderCmd                    `cmd:"" name:"render" aliases:"r" help:"Renders K8s Manifests"`

	// used internally by the CLI to render the manifests
	Synth SynthCmd `cmd:"" name:"synth" help:"Creates K8s Manifests" hidden:"true"`
}

type PulumiCommands struct {
	Up UpCmd `cmd:"" name:"up" help:"Runs"`
}

type InstallCommands struct {
	Install InstallCmd `cmd:"" name:"install" aliases:"i" help:"(Re)Installs the soloware CLI"`
}
