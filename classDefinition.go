package XcodeGo

type ClassDefinitionLanguage int

const (
	ObjectiveC ClassDefinitionLanguage = iota
	ObjectiveCPlusPlus
	CPlusPlus
)

type ClassDefinition struct {
	XCAbstractDefinition
	ClassName string
	Header    string
	Source    string
	language  ClassDefinitionLanguage
}

func NewClassDefinition(className string, language ClassDefinitionLanguage) *ClassDefinition {
	return &ClassDefinition{
		XCAbstractDefinition: XCAbstractDefinition{XCFileOperationTypeOverwrite},
		ClassName:            className,
		language:             language,
	}
}

func NewObjCClassDefinitionWithName(className string) *ClassDefinition {
	return NewClassDefinition(className, ObjectiveC)
}

func (t ClassDefinition) IsObjectiveC() bool {
	return t.language == ObjectiveC
}

func (t ClassDefinition) IsCPlusPlus() bool {
	return t.language == CPlusPlus
}

func (t ClassDefinition) IsObjectiveCPlusPlus() bool {
	return t.language == ObjectiveCPlusPlus
}

func (t ClassDefinition) HeaderFileName() string {
	if len(t.ClassName) == 0 {
		return ""
	}
	return t.ClassName + ".h"
}

func (t ClassDefinition) SourceFileName() (sourceFileName string) {
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
