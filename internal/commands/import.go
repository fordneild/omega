package commands

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/fordneild/omega/internal/globals"
	"github.com/fordneild/omega/internal/k8sapi"
	"github.com/fordneild/omega/internal/shell"
)

var wellKnownCRDs = []string{
	"applications.argoproj.io",
	"appprojects.argoproj.io",
	"providers.pkg.crossplane.io",
	"providerconfigs.aws.upbound.io",
}

type ImportResourceDefinitionsCmd struct {
	KubeconfigPath string `flag:"kubeconfig" optional:"" help:"Path to kubeconfig file" default:"./kubeconfig.json"`
}

func (cmd *ImportResourceDefinitionsCmd) Run(globals *globals.Globals) error {
	log.Infof("Import kubernetes resources definitions")

	config, err := k8sapi.GetKubeconfig(cmd.KubeconfigPath)
	if err != nil {
		return err
	}

	crds, err := k8sapi.GetCRDs(context.Background(), config, wellKnownCRDs)
	if err != nil {
		return err
	}

	err = cdk8sImport(crds)
	if err != nil {
		return err
	}

	return nil
}

func cdk8sImport(crds string) error {
	// create a temp file and write the crds to it
	tempFile, err := os.CreateTemp("", "crds.json")
	if err != nil {
		return err
	}
	defer os.Remove(tempFile.Name())
	tempFile.WriteString(crds)

	shellCmd := fmt.Sprintf("cat %s | cdk8s import --language=go --check-upgrade=false /dev/stdin", tempFile.Name())

	stdout, stderr, err := shell.Exec(shellCmd)
	if stdout != "" {
		log.Print(stdout)
	}
	if err != nil {
		if stderr != "" {
			log.Error(stderr)
		}
		return err
	}
	return nil
}
