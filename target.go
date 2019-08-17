package XcodeGo

type XcodeTarget struct {
	Key              string
	Name             string
	ProductName      string
	ProductReference string
	ProductType      string

	members           []BuildFile
	resources         []SourceFile
	buildShellScripts []BuildShellScript
}

func (x XcodeTarget) Resources() []SourceFile {
	return x.resources
}

func (x XcodeTarget) Members() []BuildFile {
	return x.members
}

func (x XcodeTarget) BuildShellScripts() []BuildShellScript {
	return x.buildShellScripts
}

func (x XcodeTarget) Configurations() map[string]XcodeBuildConfig {

}

func (x XcodeTarget) ConfigurationWithName(name string) XcodeBuildConfig {

}

func (x XcodeTarget) DefaultConfiguration() XcodeBuildConfig {

}

func (x *XcodeTarget) AddMember(member BuildFile) {

}

func (x *XcodeTarget) MakeAndAddShellScript(shellScript BuildShellScriptDefinition) {

}

func (x *XcodeTarget) RemoveShellScriptByName(name string) {

}

func (x *XcodeTarget) RemoveMembersWithKeys(keys []string) {

}

func (x *XcodeTarget) RemoveSourceWithKey(key string) {

}

func (x *XcodeTarget) RemoveSourcesWithKeys(keys []string) {

}

func (x *XcodeTarget) AddDependency(key string) {

}

func (x XcodeTarget) DuplicateWithTargetNameAndProductName(target string, product string) XcodeTarget {

}

func (x XcodeTarget) IsApplication() bool {

}
