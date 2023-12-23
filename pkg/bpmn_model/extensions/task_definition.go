package extensions

// TTaskDefinition is the BPMN XML representation of a task definition
type TTaskDefinition struct {
	TypeName string `xml:"type,attr"`
	Retries  string `xml:"retries,attr"`
}
