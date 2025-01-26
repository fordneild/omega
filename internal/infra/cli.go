package infra

type InfraCLI struct {
	Up UpCmd `cmd:"" name:"up" help:"Runs"`
}
