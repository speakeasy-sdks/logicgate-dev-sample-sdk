# WorkflowMap
(*WorkflowMap*)

## Overview

A [Workflow Map](https://help.logicgate.com/hc/en-us/articles/4402683117588) represents a relationship between two Workflows

### Available Operations

* [Create](#create) - Create a workflow map
* [Delete](#delete) - Delete a workflow map
* [Read](#read) - Retrieve a workflow map
* [ReadAll](#readall) - Retrieve workflow maps
* [Update](#update) - Update a workflow map

## Create

**Permissions:** [Build Access to parent applications](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Create a workflow map from a JSON request body.

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
    res, err := s.WorkflowMap.Create(ctx, operations.CreateWorkflowMapRequest{
        WorkflowMapAPICreateIn: shared.WorkflowMapAPICreateIn{
            From: "a1b2c3d4",
            Relationship: shared.WorkflowMapAPICreateInRelationshipOneToMany,
            To: "a1b2c3d4",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WorkflowMapAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateWorkflowMapRequest](../../models/operations/createworkflowmaprequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.CreateWorkflowMapResponse](../../models/operations/createworkflowmapresponse.md), error**


## Delete

**Permissions:** [Build Access to a parent application](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Delete a workflow map specified by the ID in the URL path.

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
    res, err := s.WorkflowMap.Delete(ctx, operations.DeleteWorkflowMapRequest{
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.DeleteWorkflowMapRequest](../../models/operations/deleteworkflowmaprequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.DeleteWorkflowMapResponse](../../models/operations/deleteworkflowmapresponse.md), error**


## Read

**Permissions:** [Build Access to a parent application](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Retrieve a workflow map specified by the ID in the URL path.

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
    res, err := s.WorkflowMap.Read(ctx, operations.ReadWorkflowMapRequest{
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WorkflowMapAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ReadWorkflowMapRequest](../../models/operations/readworkflowmaprequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.ReadWorkflowMapResponse](../../models/operations/readworkflowmapresponse.md), error**


## ReadAll

**Permissions:** [Build Access](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Retrieve a page of all workflow maps that the current user has [Build Access to a parent application](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications) to.

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
    res, err := s.WorkflowMap.ReadAll(ctx, operations.ReadAllWorkflowMapsRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.PageModelOutWorkflowMapAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ReadAllWorkflowMapsRequest](../../models/operations/readallworkflowmapsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.ReadAllWorkflowMapsResponse](../../models/operations/readallworkflowmapsresponse.md), error**


## Update

**Permissions:** [Build Access to a parent application](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Update a workflow map specified by the ID in the URL path from a JSON request body. Only present properties with non-empty values are updated.

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
    res, err := s.WorkflowMap.Update(ctx, operations.UpdateWorkflowMapRequest{
        WorkflowMapAPIUpdateIn: shared.WorkflowMapAPIUpdateIn{
            Relationship: shared.WorkflowMapAPIUpdateInRelationshipManyToMany,
        },
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WorkflowMapAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateWorkflowMapRequest](../../models/operations/updateworkflowmaprequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.UpdateWorkflowMapResponse](../../models/operations/updateworkflowmapresponse.md), error**

