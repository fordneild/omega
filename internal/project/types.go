package project

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/charmbracelet/log"
)

type Project interface {
	Synth() error
	GetId() string
	GetSubId(postfix string) *string
	GetName() string
	GetPath() string
	// Repo URL
	GetRepo() string
	WithApp(app cdk8s.App)
	GetApps() []cdk8s.App
	GetTargetRevision() *string
}

type project struct {
	id             string
	name           string
	path           string
	repo           string
	apps           []cdk8s.App
	targetRevision *string
}

var _ Project = &project{}

func NewProject(id, name, path, repo string, targetRevision *string) Project {
	proj := &project{id: id, name: name, path: path, repo: repo, apps: []cdk8s.App{}, targetRevision: targetRevision}
	return proj
}

func (p *project) Synth() error {
	log.Infof("Synthing project %s which has %d apps", p.GetId(), len(p.apps))
	for _, app := range p.apps {
		log.Infof("Synthing app %s", *app.Node().Id())
		app.Synth()
	}
	return nil
}

func (p *project) WithApp(app cdk8s.App) {
	p.apps = append(p.apps, app)
}

func (p *project) GetApps() []cdk8s.App {
	return p.apps
}

func (p *project) GetId() string {
	return p.id
}

func (p *project) GetSubId(suffix string) *string {
	return jsii.Sprintf("%s-%s", p.GetId(), suffix)
}

func (p *project) GetName() string {
	return p.name
}

func (p *project) GetPath() string {
	return p.path
}

func (p *project) GetRepo() string {
	return p.repo
}

func (p *project) GetTargetRevision() *string {
	return p.targetRevision
}
