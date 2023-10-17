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


    applicationAPICreateIn := shared.ApplicationAPICreateIn{
        Color: logicgatedevsamplesdk.String("#00a3de"),
        Icon: shared.ApplicationAPICreateInIconCubes.ToPointer(),
        Name: "Cyber Risk Management Application",
        Type: shared.ApplicationAPICreateInTypeControlsCompliance.ToPointer(),
    }

    ctx := context.Background()
    res, err := s.Application.Create(ctx, applicationAPICreateIn)
    if err != nil {
        log.Fatal(err)
    }

    if res.ApplicationAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `applicationAPICreateIn`                                                       | [shared.ApplicationAPICreateIn](../../models/shared/applicationapicreatein.md) | :heavy_check_mark:                                                             | N/A                                                                            |


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
    res, err := s.Application.Delete(ctx, id)
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
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | The unique ID of the application                      |


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
    res, err := s.Application.Read(ctx, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.ApplicationAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | The unique ID of the application                      |


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

    ctx := context.Background()
    res, err := s.Application.ReadAll(ctx, page, size)
    if err != nil {
        log.Fatal(err)
    }

    if res.PageModelOutApplicationAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                 | Type                                                                                                      | Required                                                                                                  | Description                                                                                               |
| --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                     | :heavy_check_mark:                                                                                        | The context to use for the request.                                                                       |
| `page`                                                                                                    | **int*                                                                                                    | :heavy_minus_sign:                                                                                        | The zero-indexed page number (must not be less than 0, defaults to 0)                                     |
| `size`                                                                                                    | **int*                                                                                                    | :heavy_minus_sign:                                                                                        | The size of the page and maximum number of items to be returned (must not be less than 1, defaults to 20) |


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


    applicationAPIUpdateIn := shared.ApplicationAPIUpdateIn{
        Color: logicgatedevsamplesdk.String("#00a3de"),
        Icon: shared.ApplicationAPIUpdateInIconCubes.ToPointer(),
        Live: logicgatedevsamplesdk.Bool(false),
        Name: logicgatedevsamplesdk.String("Cyber Risk Management Application"),
        RestrictBuildAccess: logicgatedevsamplesdk.Bool(false),
        Type: shared.ApplicationAPIUpdateInTypeControlsCompliance.ToPointer(),
    }

    var id string = "Van"

    ctx := context.Background()
    res, err := s.Application.Update(ctx, applicationAPIUpdateIn, id)
    if err != nil {
        log.Fatal(err)
    }

    if res.ApplicationAPIOut != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `applicationAPIUpdateIn`                                                       | [shared.ApplicationAPIUpdateIn](../../models/shared/applicationapiupdatein.md) | :heavy_check_mark:                                                             | N/A                                                                            |
| `id`                                                                           | *string*                                                                       | :heavy_check_mark:                                                             | The unique ID of the application                                               |


### Response

**[*operations.Update1Response](../../models/operations/update1response.md), error**

