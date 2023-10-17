# Step
(*Step*)

## Overview

A [Step](https://help.logicgate.com/hc/en-us/articles/4402674059668-Create-a-Step) lives in a Workflow and is configured with a set of Sections, Subsections and Fields to create a form

### Available Operations

* [Create](#create) - Create a step
* [Delete](#delete) - Delete a step
* [Read](#read) - Retrieve a step
* [ReadAll](#readall) - Retrieve steps
* [Update](#update) - Update a step

## Create

**Permissions:** [Build Access to parent application](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Create a step from a JSON request body.

### Example Usage

```go
package main

import(
	"context"
	"log"
	logicgatedevsamplesdk "github.com/speakeasy-sdks/logicgate-dev-sample-sdk"
	"github.com/speakeasy-sdks/logicgate-dev-sample-sdk/pkg/models/shared"
)

func main() {
    s := logicgatedevsamplesdk.New(
        logicgatedevsamplesdk.WithSecurity(shared.Security{
            Basic: &shared.SchemeBasic{
                Password: "",
                Username: "",
            },
        }),
    )


    stepAPICreateIn := shared.StepAPICreateIn{
        AssignableUserType: shared.StepAPICreateInAssignableUserTypeAppUsers.ToPointer(),
        EnableComments: logicgatedevsamplesdk.Bool(false),
        ExternalUserMfaRequired: logicgatedevsamplesdk.Bool(false),
        Name: "Identify Risk",
        WorkflowID: "a1b2c3d4",
        Xpos: logicgatedevsamplesdk.Int(20),
        Ypos: logicgatedevsamplesdk.Int(20),
    }

    ctx := context.Background()
    res, err := s.Step.Create(ctx, stepAPICreateIn)
    if err != nil {
        log.Fatal(err)
    }

    if res.StepAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `stepAPICreateIn`                                                | [shared.StepAPICreateIn](../../models/shared/stepapicreatein.md) | :heavy_check_mark:                                               | N/A                                                              |


### Response

**[*operations.CreateStepResponse](../../models/operations/createstepresponse.md), error**


## Delete

**Permissions:** [Build Access to parent application](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Delete a step specified by the ID in the URL path.

### Example Usage

```go
package main

import(
	"context"
	"log"
	logicgatedevsamplesdk "github.com/speakeasy-sdks/logicgate-dev-sample-sdk"
	"github.com/speakeasy-sdks/logicgate-dev-sample-sdk/pkg/models/shared"
)

func main() {
    s := logicgatedevsamplesdk.New(
        logicgatedevsamplesdk.WithSecurity(shared.Security{
            Basic: &shared.SchemeBasic{
                Password: "",
                Username: "",
            },
        }),
    )


    var id string = "program"

    ctx := context.Background()
    res, err := s.Step.Delete(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.AppActionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | The unique ID of the step                             |


### Response

**[*operations.DeleteStepResponse](../../models/operations/deletestepresponse.md), error**


## Read

**Permissions:** [Build Access to parent application](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Retrieve a step specified by the ID in the URL path.

### Example Usage

```go
package main

import(
	"context"
	"log"
	logicgatedevsamplesdk "github.com/speakeasy-sdks/logicgate-dev-sample-sdk"
	"github.com/speakeasy-sdks/logicgate-dev-sample-sdk/pkg/models/shared"
)

func main() {
    s := logicgatedevsamplesdk.New(
        logicgatedevsamplesdk.WithSecurity(shared.Security{
            Basic: &shared.SchemeBasic{
                Password: "",
                Username: "",
            },
        }),
    )


    var id string = "gadzooks"

    ctx := context.Background()
    res, err := s.Step.Read(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.StepAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | The unique ID of the step                             |


### Response

**[*operations.ReadStepResponse](../../models/operations/readstepresponse.md), error**


## ReadAll

**Permissions:** [Build Access](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Retrieve a page of all steps that the current user has [Build Access to parent application](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications) to.

### Example Usage

```go
package main

import(
	"context"
	"log"
	logicgatedevsamplesdk "github.com/speakeasy-sdks/logicgate-dev-sample-sdk"
	"github.com/speakeasy-sdks/logicgate-dev-sample-sdk/pkg/models/shared"
)

func main() {
    s := logicgatedevsamplesdk.New(
        logicgatedevsamplesdk.WithSecurity(shared.Security{
            Basic: &shared.SchemeBasic{
                Password: "",
                Username: "",
            },
        }),
    )


    var page *int = 853380

    var size *int = 87498

    var workflowID *string = "Reggae"

    ctx := context.Background()
    res, err := s.Step.ReadAll(ctx, page, size, workflowID)
    if err != nil {
        log.Fatal(err)
    }

    if res.PageModelOutStepAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                         | Type                                                                                                              | Required                                                                                                          | Description                                                                                                       |
| ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                             | :heavy_check_mark:                                                                                                | The context to use for the request.                                                                               |
| `page`                                                                                                            | **int*                                                                                                            | :heavy_minus_sign:                                                                                                | The zero-indexed page number (must not be less than 0, defaults to 0)                                             |
| `size`                                                                                                            | **int*                                                                                                            | :heavy_minus_sign:                                                                                                | The size of the page and maximum number of items to be returned (must not be less than 1, defaults to 20)         |
| `workflowID`                                                                                                      | **string*                                                                                                         | :heavy_minus_sign:                                                                                                | The unique ID of a workflow where, if provided, the response will only contain steps from the identified workflow |


### Response

**[*operations.ReadAllStepsResponse](../../models/operations/readallstepsresponse.md), error**


## Update

**Permissions:** [Build Access to parent application](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Update a step specified by the ID in the URL path from a JSON request body. Only present properties with non-empty values are updated.

### Example Usage

```go
package main

import(
	"context"
	"log"
	logicgatedevsamplesdk "github.com/speakeasy-sdks/logicgate-dev-sample-sdk"
	"github.com/speakeasy-sdks/logicgate-dev-sample-sdk/pkg/models/shared"
)

func main() {
    s := logicgatedevsamplesdk.New(
        logicgatedevsamplesdk.WithSecurity(shared.Security{
            Basic: &shared.SchemeBasic{
                Password: "",
                Username: "",
            },
        }),
    )


    stepAPIUpdateIn := shared.StepAPIUpdateIn{
        AssignableUserType: shared.StepAPIUpdateInAssignableUserTypeAppUsers.ToPointer(),
        EnableComments: logicgatedevsamplesdk.Bool(false),
        ExternalUserMfaRequired: logicgatedevsamplesdk.Bool(false),
        Name: logicgatedevsamplesdk.String("Identify Risk"),
        Type: shared.StepAPIUpdateInTypeOrigin.ToPointer(),
        Xpos: logicgatedevsamplesdk.Int(20),
        Ypos: logicgatedevsamplesdk.Int(20),
    }

    var id string = "Van"

    ctx := context.Background()
    res, err := s.Step.Update(ctx, stepAPIUpdateIn, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.StepAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `stepAPIUpdateIn`                                                | [shared.StepAPIUpdateIn](../../models/shared/stepapiupdatein.md) | :heavy_check_mark:                                               | N/A                                                              |
| `id`                                                             | *string*                                                         | :heavy_check_mark:                                               | The unique ID of the step                                        |


### Response

**[*operations.UpdateResponse](../../models/operations/updateresponse.md), error**

