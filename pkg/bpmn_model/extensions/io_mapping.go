package extensions

// TIoMapping is the BPMN XML representation of a io mapping
type TIoMapping struct {
	Source string `xml:"source,attr"`
	Target string `xml:"target,attr"`
}
