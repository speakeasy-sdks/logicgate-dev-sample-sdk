# github.com/speakeasy-sdks/logicgate-dev-sample-sdk

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://github.com/speakeasy-sdks/logicgate-dev-sample-sdk.git/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/logicgate-dev-sample-sdk/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
    
</div>

<!-- Start SDK Installation -->
# SDK Installation

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
	"github.com/speakeasy-sdks/logicgate-dev-sample-sdk/pkg/models/operations"
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

	ctx := context.Background()
	res, err := s.Application.CreateApplication(ctx, operations.CreateApplicationRequest{
		ApplicationAPICreateIn: shared.ApplicationAPICreateIn{
			Color: logicgatedevsamplesdk.String("#00a3de"),
			Icon:  shared.ApplicationAPICreateInIconCubes.ToPointer(),
			Name:  "Cyber Risk Management Application",
			Type:  shared.ApplicationAPICreateInTypeControlsCompliance.ToPointer(),
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
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
# Available Resources and Operations


## [Application](docs/sdks/application/README.md)

* [CreateApplication](docs/sdks/application/README.md#createapplication) - Create an application
* [DeleteApplication](docs/sdks/application/README.md#deleteapplication) - Delete an application
* [ReadAllApplications](docs/sdks/application/README.md#readallapplications) - Retrieve applications
* [ReadApplication](docs/sdks/application/README.md#readapplication) - Retrieve an application
* [Update1](docs/sdks/application/README.md#update1) - Update an application

## [Authentication](docs/sdks/authentication/README.md)

* [GetAPIToken](docs/sdks/authentication/README.md#getapitoken) - Create an API Access Token

## [Field](docs/sdks/field/README.md)

* [ReadAllFields](docs/sdks/field/README.md#readallfields) - Retrieve fields

## [Record](docs/sdks/record/README.md)

* [ReadAllRecords](docs/sdks/record/README.md#readallrecords) - Retrieve records

## [Step](docs/sdks/step/README.md)

* [CreateStep](docs/sdks/step/README.md#createstep) - Create a step
* [DeleteStep](docs/sdks/step/README.md#deletestep) - Delete a step
* [ReadAllSteps](docs/sdks/step/README.md#readallsteps) - Retrieve steps
* [ReadStep](docs/sdks/step/README.md#readstep) - Retrieve a step
* [Update](docs/sdks/step/README.md#update) - Update a step

## [Workflow](docs/sdks/workflow/README.md)

* [CreateWorkflow](docs/sdks/workflow/README.md#createworkflow) - Create a workflow
* [DeleteWorkflow](docs/sdks/workflow/README.md#deleteworkflow) - Delete a workflow
* [ReadAllWorkflows](docs/sdks/workflow/README.md#readallworkflows) - Retrieve workflows
* [ReadWorkflow](docs/sdks/workflow/README.md#readworkflow) - Retrieve a workflow
* [UpdateWorkflow](docs/sdks/workflow/README.md#updateworkflow) - Update a workflow

## [WorkflowMap](docs/sdks/workflowmap/README.md)

* [CreateWorkflowMap](docs/sdks/workflowmap/README.md#createworkflowmap) - Create a workflow map
* [DeleteWorkflowMap](docs/sdks/workflowmap/README.md#deleteworkflowmap) - Delete a workflow map
* [ReadAllWorkflowMaps](docs/sdks/workflowmap/README.md#readallworkflowmaps) - Retrieve workflow maps
* [ReadWorkflowMap](docs/sdks/workflowmap/README.md#readworkflowmap) - Retrieve a workflow map
* [UpdateWorkflowMap](docs/sdks/workflowmap/README.md#updateworkflowmap) - Update a workflow map
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
