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