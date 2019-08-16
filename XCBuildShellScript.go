package XcodeGo

type XCBuildShellScript struct {
	Key                                string
	Name                               string
	RunOnlyForDeploymentPostprocessing bool
	ShellScript                        string
	ShellPath                          string
	Files                              []string
	InputFiles                         []string
	OutputFiles                        []string
}
