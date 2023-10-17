# Application
(*Application*)

## Overview

An [Application](https://help.logicgate.com/hc/en-us/articles/4402674055572-Create-a-new-Application) is a collection of Workflows, Steps, and logic that collectively solve a business use case

### Available Operations

* [Create](#create) - Create an application
* [Delete](#delete) - Delete an application
* [Read](#read) - Retrieve an application
* [ReadAll](#readall) - Retrieve applications
* [Update](#update) - Update an application

## Create

**Permissions:** [Build Access](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Create an application from a JSON request body.

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
    res, err := s.Application.Create(ctx, operations.CreateApplicationRequest{
        ApplicationAPICreateIn: shared.ApplicationAPICreateIn{
            Color: logicgatedevsamplesdk.String("#00a3de"),
            Icon: shared.ApplicationAPICreateInIconCubes.ToPointer(),
            Name: "Cyber Risk Management Application",
            Type: shared.ApplicationAPICreateInTypeControlsCompliance.ToPointer(),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ApplicationAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateApplicationRequest](../../models/operations/createapplicationrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.CreateApplicationResponse](../../models/operations/createapplicationresponse.md), error**


## Delete

**Permissions:** [Build Access to application](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Delete an application specified by the ID in the URL path.

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
    res, err := s.Application.Delete(ctx, operations.DeleteApplicationRequest{
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
| `request`                                                                                  | [operations.DeleteApplicationRequest](../../models/operations/deleteapplicationrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.DeleteApplicationResponse](../../models/operations/deleteapplicationresponse.md), error**


## Read

**Permissions:** [Build Access to application](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Retrieve an application specified by the ID in the URL path.

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
    res, err := s.Application.Read(ctx, operations.ReadApplicationRequest{
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ApplicationAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ReadApplicationRequest](../../models/operations/readapplicationrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.ReadApplicationResponse](../../models/operations/readapplicationresponse.md), error**


## ReadAll

**Permissions:** [Build Access](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Retrieve a page of all applications that the current user has [Build Access](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications) to.

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
    res, err := s.Application.ReadAll(ctx, operations.ReadAllApplicationsRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.PageModelOutApplicationAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ReadAllApplicationsRequest](../../models/operations/readallapplicationsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.ReadAllApplicationsResponse](../../models/operations/readallapplicationsresponse.md), error**


## Update

**Permissions:** [Build Access to application](https://help.logicgate.com/hc/en-us/articles/4402683190164-Control-Build-Access-for-Applications)

Update an application specified by the ID in the URL path from a JSON request body. Only present properties with non-empty values are updated.

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
    res, err := s.Application.Update(ctx, operations.Update1Request{
        ApplicationAPIUpdateIn: shared.ApplicationAPIUpdateIn{
            Color: logicgatedevsamplesdk.String("#00a3de"),
            Icon: shared.ApplicationAPIUpdateInIconCubes.ToPointer(),
            Live: logicgatedevsamplesdk.Bool(false),
            Name: logicgatedevsamplesdk.String("Cyber Risk Management Application"),
            RestrictBuildAccess: logicgatedevsamplesdk.Bool(false),
            Type: shared.ApplicationAPIUpdateInTypeControlsCompliance.ToPointer(),
        },
        ID: "<ID>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ApplicationAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [operations.Update1Request](../../models/operations/update1request.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.Update1Response](../../models/operations/update1response.md), error**

