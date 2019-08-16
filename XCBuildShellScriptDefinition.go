package XcodeGo

type XCBuildShellScriptDefinition struct {
	XCAbstractDefinition
	Key                                string
	Files                              []string
	InputPaths                         []string
	OutPutPaths                        []string
	Name                               string
	RunOnlyForDeploymentPostprocessing bool
	ShellScript                        string
	ShellPath                          string
}

func NewXCBuildShellScriptDefinition(name string, runOnlyForDeploymentPostprocessing bool, shellScript string) XCBuildShellScriptDefinition {
	return XCBuildShellScriptDefinition{
		XCAbstractDefinition:               XCAbstractDefinition{XCFileOperationTypeOverwrite},
		Files:                              []string{},
		InputPaths:                         []string{},
		OutPutPaths:                        []string{},
		Name:                               name,
		RunOnlyForDeploymentPostprocessing: runOnlyForDeploymentPostprocessing,
		ShellScript:                        shellScript,
		ShellPath:                          "/bin/sh",
	}
}
