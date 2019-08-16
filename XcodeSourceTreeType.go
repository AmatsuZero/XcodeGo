package XcodeGo

type XcodeSourceTreeType int

const (
	SourceTreeSDKRoot XcodeSourceTreeType = iota
	SourceTreeGroup
)

func (t XcodeSourceTreeType) String() string {
	return [...]string{
		"SDKROOT",
		"<group>",
	}[t]
}
