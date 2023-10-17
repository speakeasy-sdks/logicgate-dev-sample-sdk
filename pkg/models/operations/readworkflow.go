// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/logicgate-dev-sample-sdk/pkg/models/shared"
	"net/http"
)

type ReadWorkflowRequest struct {
	// The unique ID of the workflow
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *ReadWorkflowRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type ReadWorkflowResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	WorkflowAPIOut *shared.WorkflowAPIOut
}

func (o *ReadWorkflowResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ReadWorkflowResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ReadWorkflowResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ReadWorkflowResponse) GetWorkflowAPIOut() *shared.WorkflowAPIOut {
	if o == nil {
		return nil
	}
	return o.WorkflowAPIOut
}
