package bpmnmodel

import "github.com/FastGeo/bpmn-model/pkg/bpmn_model/extensions"

// TDefinitions is the BPMN XML representation of a BPMN model
type TDefinitions struct {
	ID                 string     `xml:"id,attr"`              // ID 流程定义
	Name               string     `xml:"name,attr"`            // Name 流程定义名称
	TargetNamespace    string     `xml:"targetNamespace,attr"` // xml
	ExpressionLanguage string     `xml:"expressionLanguage,attr"`
	TypeLanguage       string     `xml:"typeLanguage,attr"`
	Exporter           string     `xml:"exporter,attr"`
	ExporterVersion    string     `xml:"exporterVersion,attr"`
	Process            TProcess   `xml:"process"`
	Messages           []TMessage `xml:"message"`
}

// TProcess is the BPMN XML representation of a process
type TProcess struct {
	ID                           string                    `xml:"id,attr"`
	Name                         string                    `xml:"name,attr"`
	ProcessType                  string                    `xml:"processType,attr"`
	IsClosed                     bool                      `xml:"isClosed,attr"`
	IsExecutable                 bool                      `xml:"isExecutable,attr"`
	DefinitionalCollaborationRef string                    `xml:"definitionalCollaborationRef,attr"`
	StartEvents                  []TStartEvent             `xml:"startEvent"`
	EndEvents                    []TEndEvent               `xml:"endEvent"`
	SequenceFlows                []TSequenceFlow           `xml:"sequenceFlow"`
	ServiceTasks                 []TServiceTask            `xml:"serviceTask"`
	UserTasks                    []TUserTask               `xml:"userTask"`
	ParallelGateway              []TParallelGateway        `xml:"parallelGateway"`
	ExclusiveGateway             []TExclusiveGateway       `xml:"exclusiveGateway"`
	IntermediateCatchEvent       []TIntermediateCatchEvent `xml:"intermediateCatchEvent"`
	IntermediateTrowEvent        []TIntermediateThrowEvent `xml:"intermediateThrowEvent"`
	EventBasedGateway            []TEventBasedGateway      `xml:"eventBasedGateway"`
}

// TSequenceFlow is the BPMN XML representation of a sequence flow
type TSequenceFlow struct {
	ID                  string        `xml:"id,attr"`
	Name                string        `xml:"name,attr"`
	SourceRef           string        `xml:"sourceRef,attr"`
	TargetRef           string        `xml:"targetRef,attr"`
	ConditionExpression []TExpression `xml:"conditionExpression"`
}

// TExpression is the BPMN XML representation of an expression
type TExpression struct {
	Text string `xml:",innerxml"`
}

// TStartEvent BPMN Start Event
type TStartEvent struct {
	ID                  string   `xml:"id,attr"`
	Name                string   `xml:"name,attr"`
	IsInterrupting      bool     `xml:"isInterrupting,attr"`
	ParallelMultiple    bool     `xml:"parallelMultiple,attr"`
	IncomingAssociation []string `xml:"incoming"`
	OutgoingAssociation []string `xml:"outgoing"`
}

// TEndEvent BPM End Event
type TEndEvent struct {
	ID                  string   `xml:"id,attr"`
	Name                string   `xml:"name,attr"`
	IncomingAssociation []string `xml:"incoming"`
	OutgoingAssociation []string `xml:"outgoing"`
}

// TServiceTask BPMN Service Task
type TServiceTask struct {
	ID                  string                     `xml:"id,attr"`
	Name                string                     `xml:"name,attr"`
	Default             string                     `xml:"default,attr"`
	CompletionQuantity  int                        `xml:"completionQuantity,attr"`
	IsForCompensation   bool                       `xml:"isForCompensation,attr"`
	OperationRef        string                     `xml:"operationRef,attr"`
	Implementation      string                     `xml:"implementation,attr"`
	IncomingAssociation []string                   `xml:"incoming"`
	OutgoingAssociation []string                   `xml:"outgoing"`
	Input               []extensions.TIoMapping    `xml:"extensionElements>ioMapping>input"`
	Output              []extensions.TIoMapping    `xml:"extensionElements>ioMapping>output"`
	TaskDefinition      extensions.TTaskDefinition `xml:"extensionElements>taskDefinition"`
}

// TUserTask BPMN User Task
type TUserTask struct {
	ID                   string                           `xml:"id,attr"`
	Name                 string                           `xml:"name,attr"`
	IncomingAssociation  []string                         `xml:"incoming"`
	OutgoingAssociation  []string                         `xml:"outgoing"`
	Input                []extensions.TIoMapping          `xml:"extensionElements>ioMapping>input"`
	Output               []extensions.TIoMapping          `xml:"extensionElements>ioMapping>output"`
	AssignmentDefinition extensions.TAssignmentDefinition `xml:"extensionElements>assignmentDefinition"`
}

// TParallelGateway BPMN Parallel Gateway
type TParallelGateway struct {
	ID                  string   `xml:"id,attr"`
	Name                string   `xml:"name,attr"`
	IncomingAssociation []string `xml:"incoming"`
	OutgoingAssociation []string `xml:"outgoing"`
}

// TExclusiveGateway BPMN Exclusive Gateway
type TExclusiveGateway struct {
	ID                  string   `xml:"id,attr"`
	Name                string   `xml:"name,attr"`
	IncomingAssociation []string `xml:"incoming"`
	OutgoingAssociation []string `xml:"outgoing"`
}

// TIntermediateCatchEvent BPMN Intermediate Catch Event
type TIntermediateCatchEvent struct {
	ID                     string                  `xml:"id,attr"`
	Name                   string                  `xml:"name,attr"`
	IncomingAssociation    []string                `xml:"incoming"`
	OutgoingAssociation    []string                `xml:"outgoing"`
	MessageEventDefinition TMessageEventDefinition `xml:"messageEventDefinition"`
	TimerEventDefinition   TTimerEventDefinition   `xml:"timerEventDefinition"`
	LinkEventDefinition    TLinkEventDefinition    `xml:"linkEventDefinition"`
	ParallelMultiple       bool                    `xml:"parallelMultiple"`
	Output                 []extensions.TIoMapping `xml:"extensionElements>ioMapping>output"`
}

// TIntermediateThrowEvent BPMN Intermediate Throw Event
type TIntermediateThrowEvent struct {
	ID                  string                  `xml:"id,attr"`
	Name                string                  `xml:"name,attr"`
	IncomingAssociation []string                `xml:"incoming"`
	LinkEventDefinition TLinkEventDefinition    `xml:"linkEventDefinition"`
	Output              []extensions.TIoMapping `xml:"extensionElements>ioMapping>output"`
}

// TEventBasedGateway BPMN Event Based Gateway
type TEventBasedGateway struct {
	ID                  string   `xml:"id,attr"`
	Name                string   `xml:"name,attr"`
	IncomingAssociation []string `xml:"incoming"`
	OutgoingAssociation []string `xml:"outgoing"`
}

// TMessageEventDefinition BPMN Message Event Definition
type TMessageEventDefinition struct {
	ID         string `xml:"id,attr"`
	MessageRef string `xml:"messageRef,attr"`
}

// TTimerEventDefinition BPMN Timer Event Definition
type TTimerEventDefinition struct {
	ID           string        `xml:"id,attr"`
	TimeDuration TTimeDuration `xml:"timeDuration"`
}

// TLinkEventDefinition BPMN Link Event Definition
type TLinkEventDefinition struct {
	ID   string `xml:"id,attr"`
	Name string `xml:"name,attr"`
}

// TMessage BPMN Message
type TMessage struct {
	ID   string `xml:"id,attr"`
	Name string `xml:"name,attr"`
}

// TTimeDuration BPMN Time Duration
type TTimeDuration struct {
	XMLText string `xml:",innerxml"`
}

// TInclusiveGateway BPMN Inclusive Gateway
type TInclusiveGateway struct {
	ID                  string   `xml:"id,attr"`
	Name                string   `xml:"name,attr"`
	IncomingAssociation []string `xml:"incoming"`
	OutgoingAssociation []string `xml:"outgoing"`
}
