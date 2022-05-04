package batch

import (
	"github.com/PixarV/aws-sdk-go/aws"
	"github.com/PixarV/aws-sdk-go/service/batch"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/PixarV/terraform-provider-ritt/internal/tfresource"
)

func statusComputeEnvironment(conn *batch.Batch, name string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		computeEnvironmentDetail, err := FindComputeEnvironmentDetailByName(conn, name)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return computeEnvironmentDetail, aws.StringValue(computeEnvironmentDetail.Status), nil
	}
}
