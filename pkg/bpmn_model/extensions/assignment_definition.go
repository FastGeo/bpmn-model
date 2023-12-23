package extensions

import "strings"

// TAssignmentDefinition is the BPMN XML representation of a assignment definition
type TAssignmentDefinition struct {
	Assignee        string `xml:"assignee,attr"`
	CandidateGroups string `xml:"candidateGroups,attr"`
}

// GetCandidateGroups returns the candidate groups of the assignment definition
func (ad TAssignmentDefinition) GetCandidateGroups() []string {
	groups := strings.Split(ad.CandidateGroups, ",")
	for i, group := range groups {
		groups[i] = strings.TrimSpace(group)
	}
	return groups
}
