package XcodeGo

type XCBuildFile interface {
	BecomeBuildFile()
	BuildPhase() XcodeMemberType
	BuildFileKey() string
	IsBuildFile() bool
}
