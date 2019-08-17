package XcodeGo

type BuildFile interface {
	BecomeBuildFile()
	BuildPhase() XcodeMemberType
	BuildFileKey() string
	IsBuildFile() bool
}
