package commands

type AllCommands struct {
	Cdk8sCommands
	PulumiCommands
	InstallCommands
}

type Cdk8sCommands struct {
	Import ImportResourceDefinitionsCmd `cmd:"" name:"import" help:"Imports K8s resource definitions from a cluster/url"`
	Synth  SynthCmd                     `cmd:"" name:"synth" help:"Creates K8s Manifests"`
	Sync   SyncCmd                      `cmd:"" name:"sync" help:"Syncs K8s Manifests"`
}

type PulumiCommands struct {
	Up UpCmd `cmd:"" name:"up" help:"Runs"`
}

type InstallCommands struct {
	Install InstallCmd `cmd:"" name:"install" aliases:"i" help:"(Re)Installs the soloware CLI"`
}
