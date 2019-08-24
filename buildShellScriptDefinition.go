package XcodeGo

import "./Utils"

type BuildShellScriptDefinition struct {
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

func NewBuildShellScriptDefinition(name string, runOnlyForDeploymentPostprocessing bool, shellScript string) *BuildShellScriptDefinition {
	return &BuildShellScriptDefinition{
		XCAbstractDefinition:               XCAbstractDefinition{XCFileOperationTypeOverwrite},
		Files:                              []string{},
		InputPaths:                         []string{},
		OutPutPaths:                        []string{},
		Name:                               Utils.FilterEmoji(name),
		RunOnlyForDeploymentPostprocessing: runOnlyForDeploymentPostprocessing,
		ShellScript:                        shellScript,
		ShellPath:                          "/bin/sh",
	}
}
