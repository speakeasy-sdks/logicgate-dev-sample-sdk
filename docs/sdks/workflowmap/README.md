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


    workflowMapAPICreateIn := shared.WorkflowMapAPICreateIn{
        From: "a1b2c3d4",
        Relationship: shared.WorkflowMapAPICreateInRelationshipOneToMany,
        To: "a1b2c3d4",
    }

    ctx := context.Background()
    res, err := s.WorkflowMap.Create(ctx, workflowMapAPICreateIn)
    if err != nil {
        log.Fatal(err)
    }

    if res.WorkflowMapAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `workflowMapAPICreateIn`                                                       | [shared.WorkflowMapAPICreateIn](../../models/shared/workflowmapapicreatein.md) | :heavy_check_mark:                                                             | N/A                                                                            |


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
    res, err := s.WorkflowMap.Delete(ctx, id)
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
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | The unique ID of the workflow map                     |


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
    res, err := s.WorkflowMap.Read(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.WorkflowMapAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | The unique ID of the workflow map                     |


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
    res, err := s.WorkflowMap.ReadAll(ctx, page, size, workflowID)
    if err != nil {
        log.Fatal(err)
    }

    if res.PageModelOutWorkflowMapAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                 | Type                                                                                                                      | Required                                                                                                                  | Description                                                                                                               |
| ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                                     | :heavy_check_mark:                                                                                                        | The context to use for the request.                                                                                       |
| `page`                                                                                                                    | **int*                                                                                                                    | :heavy_minus_sign:                                                                                                        | The zero-indexed page number (must not be less than 0, defaults to 0)                                                     |
| `size`                                                                                                                    | **int*                                                                                                                    | :heavy_minus_sign:                                                                                                        | The size of the page and maximum number of items to be returned (must not be less than 1, defaults to 20)                 |
| `workflowID`                                                                                                              | **string*                                                                                                                 | :heavy_minus_sign:                                                                                                        | The unique ID of a workflow where, if provided, the response will only contain workflow maps from the identified workflow |


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


    workflowMapAPIUpdateIn := shared.WorkflowMapAPIUpdateIn{
        Relationship: shared.WorkflowMapAPIUpdateInRelationshipManyToMany,
    }

    var id string = "New"

    ctx := context.Background()
    res, err := s.WorkflowMap.Update(ctx, workflowMapAPIUpdateIn, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.WorkflowMapAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `workflowMapAPIUpdateIn`                                                       | [shared.WorkflowMapAPIUpdateIn](../../models/shared/workflowmapapiupdatein.md) | :heavy_check_mark:                                                             | N/A                                                                            |
| `id`                                                                           | *string*                                                                       | :heavy_check_mark:                                                             | The unique ID of the workflow map                                              |


### Response

**[*operations.UpdateWorkflowMapResponse](../../models/operations/updateworkflowmapresponse.md), error**

