package k8sapi

import (
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func GetKubeconfig(path string) (*rest.Config, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("kubeconfig file not found at %s - please specify a valid path using --kubeconfig-path", path)
	}

	return ParseKubeconfig(path)
}

func ParseKubeconfig(path string) (*rest.Config, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve path: %w", err)
	}

	config, err := clientcmd.BuildConfigFromFlags("", absPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load kubeconfig: %w", err)
	}

	return config, nil
}
