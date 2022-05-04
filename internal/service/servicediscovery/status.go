package servicediscovery

import (
	"github.com/PixarV/aws-sdk-go/aws"
	"github.com/PixarV/aws-sdk-go/service/servicediscovery"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/PixarV/terraform-provider-ritt/internal/tfresource"
)

// StatusOperation fetches the Operation and its Status
func StatusOperation(conn *servicediscovery.ServiceDiscovery, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		output, err := FindOperationByID(conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.Status), nil
	}
}
