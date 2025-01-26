package resources

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/fordneild/omega/internal/globals"
)

type SynthCmd struct {
}

func (cmd *SynthCmd) Run(globals *globals.Globals) error {
	app := cdk8s.NewApp(nil)
	NewChart(app, "getting-started", "default", "my-app")
	app.Synth()
	return nil
}
