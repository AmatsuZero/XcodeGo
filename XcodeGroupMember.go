package XcodeGo

type XcodeGroupMember interface {
	Key() string
	DisplayName() string
	PathRelativeToProjectRoot() string
	GroupMemberType() XcodeMemberType
}
