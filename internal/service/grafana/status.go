package grafana

import (
	"github.com/PixarV/aws-sdk-go/aws"
	"github.com/PixarV/aws-sdk-go/service/managedgrafana"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/PixarV/terraform-provider-ritt/internal/tfresource"
)

func statusWorkspaceStatus(conn *managedgrafana.ManagedGrafana, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		output, err := FindWorkspaceByID(conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.Status), nil
	}
}

func statusWorkspaceSamlConfiguration(conn *managedgrafana.ManagedGrafana, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		output, err := FindSamlConfigurationByID(conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.Status), nil
	}
}
