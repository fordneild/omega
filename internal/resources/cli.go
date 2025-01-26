package resources

type ResourcesCLI struct {
	Synth SynthCmd `cmd:"" name:"synth" help:"Creates K8s Manifests"`
}
