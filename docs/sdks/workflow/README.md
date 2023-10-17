# Workflow
(*Workflow*)

## Overview

A [Workflow](https://help.logicgate.com/hc/en-us/articles/4402683108756-Create-a-new-Workflow) is a combination of Steps, Paths, Fields, and routing logic that combine to form a system in an Application

### Available Operations

* [Create](#create) - Create a workflow
* [Delete](#delete) - Delete a workflow
* [Read](#read) - Retrieve a workflow
* [ReadAll](#readall) - Retrieve workflows
* [Update](#update) - Update a workflow

## Create

**Permissions:** [Build Access to parent application](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Create a workflow from a JSON request body. The workflow will contain a Default Origin step and a Default End step.

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


    workflowAPICreateIn := shared.WorkflowAPICreateIn{
        ApplicationID: "a1b2c3d4",
        Name: "Risk Assessments",
        RecordPrefix: "Assessment",
        Xpos: logicgatedevsamplesdk.Int(20),
        Ypos: logicgatedevsamplesdk.Int(20),
    }

    ctx := context.Background()
    res, err := s.Workflow.Create(ctx, workflowAPICreateIn)
    if err != nil {
        log.Fatal(err)
    }

    if res.WorkflowAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `workflowAPICreateIn`                                                    | [shared.WorkflowAPICreateIn](../../models/shared/workflowapicreatein.md) | :heavy_check_mark:                                                       | N/A                                                                      |


### Response

**[*operations.CreateWorkflowResponse](../../models/operations/createworkflowresponse.md), error**


## Delete

**Permissions:** [Build Access to parent application](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Delete a workflow specified by the ID in the URL path.

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
    res, err := s.Workflow.Delete(ctx, id)
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
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | The unique ID of the workflow                         |


### Response

**[*operations.DeleteWorkflowResponse](../../models/operations/deleteworkflowresponse.md), error**


## Read

**Permissions:** [Build Access to parent application](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Retrieve a workflow specified by the ID in the URL path.

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
    res, err := s.Workflow.Read(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.WorkflowAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | The unique ID of the workflow                         |


### Response

**[*operations.ReadWorkflowResponse](../../models/operations/readworkflowresponse.md), error**


## ReadAll

**Permissions:** [Build Access](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Retrieve a page of all workflows that the current user has [Build Access to parent application](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications) to.

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


    var applicationID *string = "Van"

    var includeJiraWorkflows *bool = false

    var page *int = 775284

    var size *int = 45601

    ctx := context.Background()
    res, err := s.Workflow.ReadAll(ctx, applicationID, includeJiraWorkflows, page, size)
    if err != nil {
        log.Fatal(err)
    }

    if res.PageModelOutWorkflowAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `applicationID`                                                                                                                                      | **string*                                                                                                                                            | :heavy_minus_sign:                                                                                                                                   | The unique ID of a parent application where, if provided, the response will only contain workflows from the identified application                   |
| `includeJiraWorkflows`                                                                                                                               | **bool*                                                                                                                                              | :heavy_minus_sign:                                                                                                                                   | Whether [Jira workflows](https://help.logicgate.com/hc/en-us) are to be included in the response in addition to regular workflows (defaults to true) |
| `page`                                                                                                                                               | **int*                                                                                                                                               | :heavy_minus_sign:                                                                                                                                   | The zero-indexed page number (must not be less than 0, defaults to 0)                                                                                |
| `size`                                                                                                                                               | **int*                                                                                                                                               | :heavy_minus_sign:                                                                                                                                   | The size of the page and maximum number of items to be returned (must not be less than 1, defaults to 20)                                            |


### Response

**[*operations.ReadAllWorkflowsResponse](../../models/operations/readallworkflowsresponse.md), error**


## Update

**Permissions:** [Build Access to parent application](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Update a workflow specified by the ID in the URL path from a JSON request body. Only present properties with non-empty values are updated.

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


    workflowAPIUpdateIn := shared.WorkflowAPIUpdateIn{
        Name: logicgatedevsamplesdk.String("Risk Assessments"),
        RecordPrefix: logicgatedevsamplesdk.String("Assessment"),
        Xpos: logicgatedevsamplesdk.Int(20),
        Ypos: logicgatedevsamplesdk.Int(20),
    }

    var id string = "Van"

    ctx := context.Background()
    res, err := s.Workflow.Update(ctx, workflowAPIUpdateIn, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.WorkflowAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `workflowAPIUpdateIn`                                                    | [shared.WorkflowAPIUpdateIn](../../models/shared/workflowapiupdatein.md) | :heavy_check_mark:                                                       | N/A                                                                      |
| `id`                                                                     | *string*                                                                 | :heavy_check_mark:                                                       | The unique ID of the workflow                                            |


### Response

**[*operations.UpdateWorkflowResponse](../../models/operations/updateworkflowresponse.md), error**

