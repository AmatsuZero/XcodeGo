package XcodeGo

type ProjectBuildConfig struct {
	project               *Project
	Key                   string
	SpecificBuildSettings map[string]string
}

func BuildConfiuration(configs []ProjectBuildConfig, project *Project) map[string]string {

}

func NewProjectBuildConfig(project *Project, key string) ProjectBuildConfig {

}

func (p *ProjectBuildConfig) AddBuildSettings(settings map[string]string) {

}

func (p *ProjectBuildConfig) AddOrReplaceSetting(key string, setting interface{}) {

}

func (p ProjectBuildConfig) ValueForKey() interface{} {

}

type BuildConfigurationVisitor = func(map[string]interface{})

func Duplicate(buildConfigurationListKey string, project *Project, visitor BuildConfigurationVisitor) string {

}
