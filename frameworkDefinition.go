package XcodeGo

import (
	"path/filepath"
	"strings"
)

type XCFrameworkDefinition struct {
	FilePath          string
	CopyToDestination bool
	SourceTree        XcodeSourceTreeType
}

func NewFrameworkDefinition(filePath string, copyToDestination bool, sourceTree XcodeSourceTreeType) XCFrameworkDefinition {
	return XCFrameworkDefinition{
		FilePath:          filePath,
		CopyToDestination: copyToDestination,
		SourceTree:        sourceTree,
	}
}

func NewGroupTypeFrameworkDefinition(filePath string, copyToDestination bool) XCFrameworkDefinition {
	return NewFrameworkDefinition(filePath, copyToDestination, SourceTreeGroup)
}

func (t XCFrameworkDefinition) FileName() (name string) {
	name = filepath.Base(t.FilePath)
	name = strings.ReplaceAll(name, "/", "")
	return
}
