package XcodeGo

import "fmt"

type XcodeGroup struct {
	Alias                string
	PathRelativeToParent string
	Children             []XcodeGroupMember

	memberType         XcodeMemberType
	project            *Project
	buildFileKey       string
	isBuildFile        bool
	fileOperationQueue FileOperationQueue
	key                string
}

func NewGroup(project *Project, key string, alias string, path string, children []XcodeGroupMember, groupType XcodeMemberType) XcodeGroup {
	if groupType != PBXGroupType && groupType != PBXVariantGroupType {
		panic(fmt.Errorf("invalid group type"))
	}
	return XcodeGroup{
		Alias:                alias,
		key:                  key,
		PathRelativeToParent: path,
		Children:             children,
		memberType:           groupType,
	}
}

func NewPBXGroup(project *Project, key string, alias string, path string, children []XcodeGroupMember) XcodeGroup {
	return NewGroup(project, key, alias, path, children, PBXGroupType)
}

func (x XcodeGroup) BecomeBuildFile() {
	panic("implement me")
}

func (x XcodeGroup) BuildPhase() XcodeMemberType {
	panic("implement me")
}

func (x XcodeGroup) BuildFileKey() string {
	panic("implement me")
}

func (x XcodeGroup) IsBuildFile() bool {
	panic("implement me")
}

func (x XcodeGroup) DisplayName() string {
	panic("implement me")
}

func (x XcodeGroup) GroupMemberType() XcodeMemberType {
	panic("implement me")
}

func (x XcodeGroup) Key() string {
	panic("implement me")
}

func (x XcodeGroup) PathRelativeToProjectRoot() string {
	panic("implement me")
}

func (g *XcodeGroup) RemoveFromParentGroup() {

}

func (g *XcodeGroup) RemoveParentDeletingChildren(deleteChildren bool) {

}

func (g XcodeGroup) ParentGroup() XcodeGroup {

}

func (g XcodeGroup) IsRootGroup() bool {

}
