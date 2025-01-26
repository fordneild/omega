package main

import (
	"github.com/fordneild/omega/internal/infra"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		err := infra.NewCluster().Run(ctx)
		return err
	})
}
