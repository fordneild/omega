package argoprojio


type ApplicationSetSpecGeneratorsClustersTemplateSpecSourcesKustomizePatchesTarget struct {
	AnnotationSelector *string `field:"optional" json:"annotationSelector" yaml:"annotationSelector"`
	Group *string `field:"optional" json:"group" yaml:"group"`
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	LabelSelector *string `field:"optional" json:"labelSelector" yaml:"labelSelector"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	Version *string `field:"optional" json:"version" yaml:"version"`
}

