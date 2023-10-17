// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// WorkflowMapAPICreateInRelationship - The type of the relationship between workflows
type WorkflowMapAPICreateInRelationship string

const (
	WorkflowMapAPICreateInRelationshipOneToOne   WorkflowMapAPICreateInRelationship = "OneToOne"
	WorkflowMapAPICreateInRelationshipOneToMany  WorkflowMapAPICreateInRelationship = "OneToMany"
	WorkflowMapAPICreateInRelationshipManyToOne  WorkflowMapAPICreateInRelationship = "ManyToOne"
	WorkflowMapAPICreateInRelationshipManyToMany WorkflowMapAPICreateInRelationship = "ManyToMany"
)

func (e WorkflowMapAPICreateInRelationship) ToPointer() *WorkflowMapAPICreateInRelationship {
	return &e
}

func (e *WorkflowMapAPICreateInRelationship) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OneToOne":
		fallthrough
	case "OneToMany":
		fallthrough
	case "ManyToOne":
		fallthrough
	case "ManyToMany":
		*e = WorkflowMapAPICreateInRelationship(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WorkflowMapAPICreateInRelationship: %v", v)
	}
}

// WorkflowMapAPICreateIn - Workflow Map (Create)
type WorkflowMapAPICreateIn struct {
	// The unique ID of the source workflow of the workflow map relationship
	From string `json:"from"`
	// The type of the relationship between workflows
	Relationship WorkflowMapAPICreateInRelationship `json:"relationship"`
	// The unique ID of the destination workflow of the workflow map relationship
	To string `json:"to"`
}

func (o *WorkflowMapAPICreateIn) GetFrom() string {
	if o == nil {
		return ""
	}
	return o.From
}

func (o *WorkflowMapAPICreateIn) GetRelationship() WorkflowMapAPICreateInRelationship {
	if o == nil {
		return WorkflowMapAPICreateInRelationship("")
	}
	return o.Relationship
}

func (o *WorkflowMapAPICreateIn) GetTo() string {
	if o == nil {
		return ""
	}
	return o.To
}