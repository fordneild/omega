package install

type InstallCommandLine struct {
	Install InstallCmd `cmd:"" name:"install" aliases:"i" help:"(Re)Installs the soloware CLI"`
}
