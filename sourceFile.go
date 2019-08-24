package XcodeGo

import (
	"./Utils"
	"fmt"
	"path/filepath"
)

type SourceFile struct {
	SourceTree   string
	project      *Project
	isBuildFile  string
	buildFileKey string
	name         string
	path         string
	sourceTree   string
	fileType     XcodeSourceFileType
}

func NewSourceFile(project *Project, key string, fileType XcodeSourceFileType, name string, sourceTree string, path string) *SourceFile {
	return &SourceFile{
		SourceTree:   sourceTree,
		path:         path,
		project:      project,
		isBuildFile:  "",
		buildFileKey: key,
		name:         name,
		sourceTree:   sourceTree,
		fileType:     fileType,
	}
}

func (s SourceFile) Type() XcodeSourceFileType {
	return s.fileType
}

func (s *SourceFile) SetName(name string) {
	s.name = name
	s.setValue(name, "name")
}

func (s SourceFile) Name() string {
	return s.name
}

func (s SourceFile) Path() string {
	return s.path
}

func (s *SourceFile) SetPath(path string) {
	s.path = path
	s.setValue(path, "path")
}

func (s *SourceFile) BecomeBuildFile() {
	if s.IsBuildFile() {
		return
	}
	if s.CanBecomeBuildFile() {
		sourceBuildFile := map[string]interface{}{
			"isa":     PBXBuildFileType,
			"fileRef": s.Key(),
		}
		keyBuilder := Utils.Unique()
		buildFileKey := keyBuilder.Build()
		objs := *s.project.Objects()
		objs[buildFileKey] = sourceBuildFile
	} else if s.fileType == Framework {
		panic("Add Framework to target not support yet")
	} else {
		panic(fmt.Sprintf("Project file of type %v cant's become a build file", s.buildFileKey))
	}

}

func (s SourceFile) CanBecomeBuildFile() bool {
	switch s.fileType {
	case SourceCodeCPlusPlus, SourceCodeObjCPlusPlus, SourceCodeObjc, XibFile, Framework, ImageResourcePNG,
		HTML, Bundle, Archive, AssetCatalog, SourceCodeSwift, PropertyList, LocalizableStrings:
		return true
	default:
		return false
	}
}

func (s *SourceFile) RemoveBuildFile() {
	if !s.IsBuildFile() {
		return
	}
	file := s.project.FileWithName(s.path)
	for k, v := range *s.project.Objects() {
		if fileRef, ok := v["fileRef"]; ok {
			if ref, ok := fileRef.(string); ok && ref == file.Key() {
				s.project.RemoveObjectWithKey(k)
			}
		}

	}
}

func (s *SourceFile) SetCompileFlags(value string) {
	objs := *s.project.Objects()
	for k, v := range objs {
		isa, ok := v["isa"]
		if !ok {
			continue
		}
		t, ok := isa.(XcodeMemberType)
		if !ok || t != PBXBuildFileType {
			continue
		}
		r, ok := v["fileRef"]
		if !ok {
			continue
		}
		ref, ok := r.(string)
		if !ok || ref != s.Key() {
			continue
		}
		replacement := v
		replacement["settings"] = map[string]interface{}{
			"COMPILER_FLAGS": value,
		}
		delete(objs, k)
		objs[k] = replacement
	}
}

func (s *SourceFile) SetWeakReference() {
	s.addAttribute("Weak")
}

func (s SourceFile) BuildPhase() XcodeMemberType {
	switch s.fileType {
	case SourceCodeObjc, SourceCodeObjCPlusPlus, SourceCodeCPlusPlus, XibFile, SourceCodeSwift:
		return PBXSourcesBuildPhaseType
	case Framework, Archive:
		return PBXFrameworksBuildPhaseType
	case ImageResourcePNG, HTML, Bundle, AssetCatalog, PropertyList, LocalizableStrings:
		return PBXResourcesBuildPhaseType
	default:
		return PBXNilType
	}
}

func (s *SourceFile) BuildFileKey() string {
	if len(s.buildFileKey) > 0 {
		return s.buildFileKey
	}
	for _, obj := range *s.project.Objects() {
		if val, ok := obj["isa"]; ok {
			if v, ok := val.(XcodeMemberType); ok && v == PBXBuildFileType {
				if ref, ok := obj["fileRef"]; ok {
					if r, ok := ref.(string); ok && r == s.Key() {
						s.buildFileKey = r
						return r
					}
				}
			}
		}
	}
	return s.buildFileKey
}

func (s *SourceFile) IsBuildFile() bool {
	if s.CanBecomeBuildFile() && len(s.isBuildFile) == 0 {
		for _, obj := range *s.project.Objects() {
			if val, ok := obj["isa"]; ok {
				if v, ok := val.(XcodeMemberType); ok && v == PBXBuildFileType {
					if ref, ok := obj["fileRef"]; ok {
						if r, ok := ref.(string); ok && r == s.Key() {
							s.isBuildFile = "YES"
							return true
						}
					}
				}
			}
		}
	}
	return s.isBuildFile == "YES"
}

func (s SourceFile) Key() string {
	return s.buildFileKey
}

func (s SourceFile) DisplayName() string {
	return s.name
}

func (s SourceFile) PathRelativeToProjectRoot() (result string) {
	result = s.project.GroupForGroupMemberWithKey(s.buildFileKey).PathRelativeToParent
	result = filepath.Join(result, s.name)
	return
}

func (s SourceFile) GroupMemberType() XcodeMemberType {
	return PBXFileReferenceType
}

func (s *SourceFile) setValue(val interface{}, key string) {
	raw := *s.project.Objects()
	if obj, ok := raw[key]; ok {
		obj[key] = val
		raw[key] = obj
	} else {
		panic("Project item with key " + key + "not found")
	}
}

func (s *SourceFile) addAttribute(attribute string) {
	objs := *s.project.Objects()
	for k, v := range objs {
		isa, ok := v["isa"]
		if !ok {
			continue
		}
		t, ok := isa.(XcodeMemberType)
		if !ok || t != PBXBuildFileType {
			continue
		}
		r, ok := v["fileRef"]
		if !ok {
			continue
		}
		ref, ok := r.(string)
		if !ok || ref != s.Key() {
			continue
		}
		replacement := v
		if _, ok := replacement["settings"]; !ok {
			replacement["settings"] = map[string]interface{}{
				"ATTRIBUTES": []string{attribute},
			}
		} else if settings, ok := replacement["settings"].(map[string]interface{}); ok {
			if _, ok := settings["ATTRIBUTES"]; !ok {
				settings["ATTRIBUTES"] = []string{attribute}
			} else if attributes, ok := settings["ATTRIBUTES"].([]string); ok {
				if ret, _ := Utils.Contain(attributes, attribute); ret {
					return
				}
				attributes = append(attributes, attribute)
				settings["ATTRIBUTES"] = attributes
			}
			replacement["settings"] = settings
		}
		delete(objs, k)
		objs[k] = replacement
	}
}
