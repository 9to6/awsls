// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListGlobalacceleratorAccelerator(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Globalacceleratorconn.ListAcceleratorsRequest(&globalaccelerator.ListAcceleratorsInput{})

	var result []terraform.Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Accelerators) > 0 {

		for _, r := range resp.Accelerators {

			t := *r.CreatedTime
			result = append(result, terraform.Resource{
				Type:      "aws_globalaccelerator_accelerator",
				ID:        *r.AcceleratorArn,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,

				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
