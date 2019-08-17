package XcodeGo

type XcodeMemberType int

const (
	PBXNilType XcodeMemberType = iota
	PBXBuildFileType
	PBXContainerItemProxyType
	PBXCopyFilesBuildPhaseType
	PBXFileReferenceType
	PBXFrameworksBuildPhaseType
	PBXGroupType
	PBXNativeTargetType
	PBXProjectType
	PBXReferenceProxyType
	PBXResourcesBuildPhaseType
	PBXSourcesBuildPhaseType
	PBXTargetDependencyType
	PBXVariantGroupType
	XCBuildConfigurationType
	XCConfigurationListType
	PBXShellScriptBuildPhase
	XCVersionGroupType
)

var memberTypeReference []string

func init() {
	memberTypeReference = append(memberTypeReference,
		"PBXNilType",
		"PBXBuildFile",
		"PBXContainerItemProxy",
		"PBXCopyFilesBuildPhase",
		"PBXFileReference",
		"PBXFrameworksBuildPhase",
		"PBXGroup",
		"PBXNativeTarget",
		"PBXProject",
		"PBXReferenceProxy",
		"PBXResourcesBuildPhase",
		"PBXShellScriptBuildPhase",
		"PBXSourcesBuildPhase",
		"PBXTargetDependency",
		"PBXVariantGroup",
		"XCBuildConfiguration",
		"XCConfigurationList",
		"XCVersionGroup")
}

func (t XcodeMemberType) String() string {
	return memberTypeReference[t]
}
