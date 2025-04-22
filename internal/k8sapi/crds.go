package k8sapi

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/fordneild/omega/internal/shell"
	"k8s.io/client-go/rest"
)

func GetCRDs(ctx context.Context, config *rest.Config, crdNames []string) (string, error) {
	jqFilter := ""
	filters := make([]string, len(crdNames))
	for i, crdName := range crdNames {
		filters[i] = fmt.Sprintf(".metadata.name == \"%s\"", crdName)
	}
	jqFilter = strings.Join(filters, " or ")
	shellCmd := fmt.Sprintf("kubectl get crds -o json | jq ' .items |= map(select(%s))'", jqFilter)
	stdout, stderr, err := shell.Exec(shellCmd)
	if err != nil {
		if stderr != "" {
			log.Error(stderr)
		}
		return "", err
	}

	return stdout, nil
}
