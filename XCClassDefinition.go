package XcodeGo

type ClassDefinitionLanguage int

const (
	ObjectiveC ClassDefinitionLanguage = iota
	ObjectiveCPlusPlus
	CPlusPlus
)

type XCClassDefinition struct {
	XCAbstractDefinition
	ClassName string
	Header    string
	Source    string
	language  ClassDefinitionLanguage
}

func NewClassDefinitionWithName(className string) XCClassDefinition {
	return XCClassDefinition{
		XCAbstractDefinition: XCAbstractDefinition{XCFileOperationTypeOverwrite},
		ClassName:            className,
		language:             ObjectiveC,
	}
}

func NewClassDefinition(className string, language ClassDefinitionLanguage) XCClassDefinition {
	return XCClassDefinition{
		XCAbstractDefinition: XCAbstractDefinition{XCFileOperationTypeOverwrite},
		ClassName:            className,
		language:             language,
	}
}

func (t XCClassDefinition) IsObjectiveC() bool {
	return t.language == ObjectiveC
}

func (t XCClassDefinition) IsCPlusPlus() bool {
	return t.language == CPlusPlus
}

func (t XCClassDefinition) IsObjectiveCPlusPlus() bool {
	return t.language == ObjectiveCPlusPlus
}

func (t XCClassDefinition) HeaderFileName() string {
	if len(t.ClassName) == 0 {
		return ""
	}
	return t.ClassName + ".h"
}

func (t XCClassDefinition) SourceFileName() (sourceFileName string) {
	if t.IsObjectiveC() {
		sourceFileName = t.ClassName + ".m"
	} else if t.IsObjectiveCPlusPlus() {
		sourceFileName = t.ClassName + ".mm"
	} else if t.IsCPlusPlus() {
		sourceFileName = t.ClassName + ".cpp"
	} else {
		sourceFileName = ""
	}
	return
}
