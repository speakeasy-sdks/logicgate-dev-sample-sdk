# github.com/speakeasy-sdks/logicgate-dev-sample-sdk

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://github.com/speakeasy-sdks/logicgate-dev-sample-sdk.git/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/logicgate-dev-sample-sdk/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
    
</div>

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/logicgate-dev-sample-sdk
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	logicgatedevsamplesdk "github.com/speakeasy-sdks/logicgate-dev-sample-sdk"
	"github.com/speakeasy-sdks/logicgate-dev-sample-sdk/pkg/models/shared"
	"log"
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
		Icon:  shared.ApplicationAPICreateInIconCubes.ToPointer(),
		Name:  "Cyber Risk Management Application",
		Type:  shared.ApplicationAPICreateInTypeControlsCompliance.ToPointer(),
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
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Application](docs/sdks/application/README.md)

* [Create](docs/sdks/application/README.md#create) - Create an application
* [Delete](docs/sdks/application/README.md#delete) - Delete an application
* [Read](docs/sdks/application/README.md#read) - Retrieve an application
* [ReadAll](docs/sdks/application/README.md#readall) - Retrieve applications
* [Update](docs/sdks/application/README.md#update) - Update an application

### [Authentication](docs/sdks/authentication/README.md)

* [GetAPIToken](docs/sdks/authentication/README.md#getapitoken) - Create an API Access Token

### [Field](docs/sdks/field/README.md)

* [ReadAll](docs/sdks/field/README.md#readall) - Retrieve fields

### [Record](docs/sdks/record/README.md)

* [ReadAll](docs/sdks/record/README.md#readall) - Retrieve records

### [Step](docs/sdks/step/README.md)

* [Create](docs/sdks/step/README.md#create) - Create a step
* [Delete](docs/sdks/step/README.md#delete) - Delete a step
* [Read](docs/sdks/step/README.md#read) - Retrieve a step
* [ReadAll](docs/sdks/step/README.md#readall) - Retrieve steps
* [Update](docs/sdks/step/README.md#update) - Update a step

### [Workflow](docs/sdks/workflow/README.md)

* [Create](docs/sdks/workflow/README.md#create) - Create a workflow
* [Delete](docs/sdks/workflow/README.md#delete) - Delete a workflow
* [Read](docs/sdks/workflow/README.md#read) - Retrieve a workflow
* [ReadAll](docs/sdks/workflow/README.md#readall) - Retrieve workflows
* [Update](docs/sdks/workflow/README.md#update) - Update a workflow

### [WorkflowMap](docs/sdks/workflowmap/README.md)

* [Create](docs/sdks/workflowmap/README.md#create) - Create a workflow map
* [Delete](docs/sdks/workflowmap/README.md#delete) - Delete a workflow map
* [Read](docs/sdks/workflowmap/README.md#read) - Retrieve a workflow map
* [ReadAll](docs/sdks/workflowmap/README.md#readall) - Retrieve workflow maps
* [Update](docs/sdks/workflowmap/README.md#update) - Update a workflow map
<!-- End SDK Available Operations -->

<!-- Start Dev Containers -->

<!-- End Dev Containers -->

<!-- Start Go Types -->

<!-- End Go Types -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
