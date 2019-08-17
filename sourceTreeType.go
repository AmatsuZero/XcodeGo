package XcodeGo

type XcodeSourceTreeType int

const (
	SourceTreeSDKRoot XcodeSourceTreeType = iota
	SourceTreeGroup
)

var sourceTreeTypeReference []string

func init() {
	sourceTreeTypeReference = []string{
		"SDKROOT",
		"<group>",
	}
}

func (t XcodeSourceTreeType) String() string {
	return sourceTreeTypeReference[t]
}
