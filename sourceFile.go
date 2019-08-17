package XcodeGo

type SourceFile struct {
	Name       string
	SourceTree string
	Path       string
}

func (s SourceFile) BecomeBuildFile() {
	panic("implement me")
}

func (s SourceFile) CanBecomeBuildFile() bool {

}

func (s SourceFile) RemoveBuildFile() {

}

func (s SourceFile) SetCompileFlags(value string) {

}

func (s SourceFile) SetWeakReference() {

}

func (s SourceFile) BuildPhase() XcodeMemberType {
	panic("implement me")
}

func (s SourceFile) BuildFileKey() string {
	panic("implement me")
}

func (s SourceFile) IsBuildFile() bool {
	panic("implement me")
}

func (s SourceFile) Key() string {
	panic("implement me")
}

func (s SourceFile) DisplayName() string {
	panic("implement me")
}

func (s SourceFile) PathRelativeToProjectRoot() string {
	panic("implement me")
}

func (s SourceFile) GroupMemberType() XcodeMemberType {
	panic("implement me")
}
