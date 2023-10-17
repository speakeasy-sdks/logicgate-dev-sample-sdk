// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package logicgatedevsamplesdk

import (
	"bytes"
	"context"
	"fmt"
	"github.com/speakeasy-sdks/logicgate-dev-sample-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/logicgate-dev-sample-sdk/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/logicgate-dev-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/logicgate-dev-sample-sdk/pkg/utils"
	"io"
	"net/http"
	"strings"
)

// authentication - Getting Started: How to create an [API Access Token](https://www.logicgate.com/developer/risk-cloud-api-authentication/) to begin integrating with the Risk Cloud API
type authentication struct {
	sdkConfiguration sdkConfiguration
}

func newAuthentication(sdkConfig sdkConfiguration) *authentication {
	return &authentication{
		sdkConfiguration: sdkConfig,
	}
}

// GetAPIToken - Create an API Access Token
// **Permissions:** Authenticated User
//
// Generates a new, expiring access token from the provided Client and Secret keys.
func (s *authentication) GetAPIToken(ctx context.Context, security operations.GetAPITokenSecurity) (*operations.GetAPITokenResponse, error) {
	request := operations.GetAPITokenRequest{}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/api/v1/account/token"

	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	utils.PopulateHeaders(ctx, req, request)

	client := utils.ConfigureSecurityClient(s.sdkConfiguration.DefaultClient, withSecurity(security))

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetAPITokenResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.LegacyAPITokenOut
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.LegacyAPITokenOut = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}
