package XcodeGo

import "path"

type XcodeSourceFileType int

const (
	FileTypeNil XcodeSourceFileType = iota
	Framework
	PropertyList
	SourceCodeHeader
	SourceCodeObjc
	SourceCodeObjCPlusPlus
	SourceCodeCPlusPlus
	XibFile
	ImageResourcePNG
	Bundle
	Archive
	HTML
	TEXT
	XcodeProject
	Folder
	AssetCatalog
	SourceCodeSwift
	Application
	Playground
	ShellScript
	Markdown
	XMLPropertyList
	Storyboard
	XCConfig
	XCDataModel
	LocalizableStrings
)

var xcfileReferenceTypes = map[string]XcodeSourceFileType{
	"sourcecode.c.h":              SourceCodeHeader,
	"sourcecode.c.objc":           SourceCodeObjc,
	"wrapper.framework":           Framework,
	"text.plist.strings":          PropertyList,
	"sourcecode.cpp.objcpp":       SourceCodeObjCPlusPlus,
	"sourcecode.cpp.cpp":          SourceCodeCPlusPlus,
	"file.xib":                    XibFile,
	"image.png":                   ImageResourcePNG,
	"wrapper.cfbundle":            Bundle,
	"archive.ar":                  Archive,
	"text.html":                   HTML,
	"text":                        TEXT,
	"wrapper.pb-project":          XcodeProject,
	"folder":                      Folder,
	"folder.assetcatalog":         AssetCatalog,
	"sourcecode.swift":            SourceCodeSwift,
	"wrapper.application":         Application,
	"file.playground":             Playground,
	"text.script.sh":              ShellScript,
	"net.daringfireball.markdown": Markdown,
	"text.plist.xml":              XMLPropertyList,
	"file.storyboard":             Storyboard,
	"text.xcconfig":               XCConfig,
	"wrapper.xcconfig":            XCConfig,
	"wrapper.xcdatamodel":         XCDataModel,
	"file.strings":                LocalizableStrings,
}

func SourceFileTypeeFromStringRepresentation(presentation string) XcodeSourceFileType {
	if fileType, ok := xcfileReferenceTypes[presentation]; ok {
		return fileType
	} else {
		return FileTypeNil
	}
}

func SourceFileTypeFromFileName(fileName string) XcodeSourceFileType {
	switch path.Ext(fileName) {
	case ".h":
		fallthrough
	case ".hh":
		fallthrough
	case ".hpp":
		fallthrough
	case ".hxx":
		return SourceCodeHeader
	case ".c":
		fallthrough
	case ".m":
		return SourceCodeObjc
	case ".mm":
		return SourceCodeObjCPlusPlus
	case ".cpp":
		return SourceCodeCPlusPlus
	case ".swift":
		return SourceCodeSwift
	case ".xcdatamodel":
		return XCDataModel
	case ".strings":
		return LocalizableStrings
	case ".plist":
		return PropertyList
	default:
		return FileTypeNil
	}
}
