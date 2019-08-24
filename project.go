package XcodeGo

type Project struct {
	FileOperationQueue *FileOperationQueue
	filePath           string
	dataSource         map[string]string
}

func NewProject(filePath string) *Project {
	return &Project{
		filePath: filePath,
	}
}

func (p Project) Files() []*SourceFile {

}

func (p Project) GroupMember(key string) *XcodeGroupMember {

}

func (p Project) FileWithKey(key string) *SourceFile {

}

func (p Project) FileWithName(name string) *SourceFile {

}

func (p Project) HeaderFiles() []*SourceFile {

}

func (p Project) SourceFiles() []*SourceFile {

}

func (p Project) XibFiles() []*SourceFile {

}

func (p Project) FilePath() string {

}

func (p Project) Groups() []*XcodeGroup {

}

func (p Project) RootGroup() *XcodeGroup {

}

func (p Project) MainGroup() *XcodeGroup {

}

func (p Project) GroupWithKey(key string) *XcodeGroup {

}

func (p Project) GroupWithDisplayName(key string) *XcodeGroup {

}

func (p Project) GroupWithPathFromRoot(path string) *XcodeGroup {

}

func (p Project) GroupForGroupMemberWithKey(key string) *XcodeGroup {

}

func (p Project) GroupWithSourceFile(sourceFile SourceFile) *SourceFile {

}

func (p *Project) PruneEmptyGroups() {

}

func (p Project) VersionGroups() []*VersionGroup {

}

func (p Project) VersionGroupWithKey(key string) *VersionGroup {

}

func (p Project) VersionGroupWithName(name string) *VersionGroup {

}

func (p Project) Targets() []*XcodeTarget {

}

func (p Project) TargetWithName(name string) *XcodeTarget {

}

func (p Project) ApplicationTargets() []*XcodeTarget {

}

func (p Project) Configurations() map[string]*ProjectBuildConfig {

}

func (p Project) ConfigurationWithName(name string) *ProjectBuildConfig {

}

func (p Project) DefaultConfiguration() *ProjectBuildConfig {

}

func (p *Project) RemoveObjectWithKey(key string) {

}

func (p *Project) Save() {

}

func (p *Project) Objects() *map[string]map[string]interface{} {

}

func (p Project) DataStore() map[string][]byte {

}

func (p *Project) DropCache() {

}
