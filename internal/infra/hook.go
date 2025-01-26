package infra

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

type PulumiHook interface {
	Run(ctx *pulumi.Context) error
}
