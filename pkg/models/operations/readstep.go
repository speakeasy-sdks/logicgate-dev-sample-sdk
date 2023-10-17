// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/logicgate-dev-sample-sdk/pkg/models/shared"
	"net/http"
)

type ReadStepRequest struct {
	// The unique ID of the step
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *ReadStepRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type ReadStepResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	StepAPIOut *shared.StepAPIOut
}

func (o *ReadStepResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ReadStepResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ReadStepResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ReadStepResponse) GetStepAPIOut() *shared.StepAPIOut {
	if o == nil {
		return nil
	}
	return o.StepAPIOut
}
