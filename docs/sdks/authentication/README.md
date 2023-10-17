# Authentication
(*Authentication*)

## Overview

Getting Started: How to create an [API Access Token](https://www.logicgate.com/developer/risk-cloud-api-authentication/) to begin integrating with the Risk Cloud API

### Available Operations

* [GetAPIToken](#getapitoken) - Create an API Access Token

## GetAPIToken

**Permissions:** Authenticated User

Generates a new, expiring access token from the provided Client and Secret keys.

### Example Usage

```go
package main

import(
	"context"
	"log"
	logicgatedevsamplesdk "github.com/speakeasy-sdks/logicgate-dev-sample-sdk"
	"github.com/speakeasy-sdks/logicgate-dev-sample-sdk/pkg/models/operations"
)

func main() {
    s := logicgatedevsamplesdk.New()


    operationSecurity := operations.GetAPITokenSecurity{
            Password: "",
            Username: "",
        }

    ctx := context.Background()
    res, err := s.Authentication.GetAPIToken(ctx, operations.GetAPITokenRequest{}, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.LegacyAPITokenOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetAPITokenRequest](../../models/operations/getapitokenrequest.md)   | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.GetAPITokenSecurity](../../models/operations/getapitokensecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


### Response

**[*operations.GetAPITokenResponse](../../models/operations/getapitokenresponse.md), error**

