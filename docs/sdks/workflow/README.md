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
	"github.com/speakeasy-sdks/logicgate-dev-sample-sdk/pkg/models/operations"
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

    ctx := context.Background()
    res, err := s.Workflow.Create(ctx, operations.CreateWorkflowRequest{
        WorkflowAPICreateIn: shared.WorkflowAPICreateIn{
            ApplicationID: "a1b2c3d4",
            Name: "Risk Assessments",
            RecordPrefix: "Assessment",
            Xpos: logicgatedevsamplesdk.Int(20),
            Ypos: logicgatedevsamplesdk.Int(20),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WorkflowAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateWorkflowRequest](../../models/operations/createworkflowrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


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
	"github.com/speakeasy-sdks/logicgate-dev-sample-sdk/pkg/models/operations"
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

    ctx := context.Background()
    res, err := s.Workflow.Delete(ctx, operations.DeleteWorkflowRequest{
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AppActionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeleteWorkflowRequest](../../models/operations/deleteworkflowrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


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
	"github.com/speakeasy-sdks/logicgate-dev-sample-sdk/pkg/models/operations"
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

    ctx := context.Background()
    res, err := s.Workflow.Read(ctx, operations.ReadWorkflowRequest{
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WorkflowAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ReadWorkflowRequest](../../models/operations/readworkflowrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


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
	"github.com/speakeasy-sdks/logicgate-dev-sample-sdk/pkg/models/operations"
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

    ctx := context.Background()
    res, err := s.Workflow.ReadAll(ctx, operations.ReadAllWorkflowsRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.PageModelOutWorkflowAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ReadAllWorkflowsRequest](../../models/operations/readallworkflowsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


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
	"github.com/speakeasy-sdks/logicgate-dev-sample-sdk/pkg/models/operations"
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

    ctx := context.Background()
    res, err := s.Workflow.Update(ctx, operations.UpdateWorkflowRequest{
        WorkflowAPIUpdateIn: shared.WorkflowAPIUpdateIn{
            Name: logicgatedevsamplesdk.String("Risk Assessments"),
            RecordPrefix: logicgatedevsamplesdk.String("Assessment"),
            Xpos: logicgatedevsamplesdk.Int(20),
            Ypos: logicgatedevsamplesdk.Int(20),
        },
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WorkflowAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateWorkflowRequest](../../models/operations/updateworkflowrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.UpdateWorkflowResponse](../../models/operations/updateworkflowresponse.md), error**

