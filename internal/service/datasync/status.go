package datasync

import (
	"github.com/PixarV/aws-sdk-go/aws"
	"github.com/PixarV/aws-sdk-go/service/datasync"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/PixarV/terraform-provider-ritt/internal/tfresource"
)

const (
	agentStatusReady = "ready"
)

func statusAgent(conn *datasync.DataSync, arn string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		output, err := FindAgentByARN(conn, arn)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, agentStatusReady, nil
	}
}

func statusTask(conn *datasync.DataSync, arn string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		output, err := FindTaskByARN(conn, arn)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.Status), nil
	}
}
