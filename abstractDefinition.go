package XcodeGo

type XCFileOperationType int

const (
	XCFileOperationTypeOverwrite XCFileOperationType = iota
	XCFileOperationTypeAcceptExisting
	XCFileOperationTypeReferenceOnly
)

type XCAbstractDefinition struct {
	FileOperationType XCFileOperationType
}
