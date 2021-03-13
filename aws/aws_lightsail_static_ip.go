// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListLightsailStaticIp(ctx context.Context, client *aws.Client) ([]terraform.Resource, error) {
	var result []terraform.Resource

	resp, err := client.Lightsailconn.GetStaticIps(ctx, &lightsail.GetStaticIpsInput{})
	if err != nil {
		return nil, err
	}

	if len(resp.StaticIps) > 0 {

		for _, r := range resp.StaticIps {

			result = append(result, terraform.Resource{
				Type:      "aws_lightsail_static_ip",
				ID:        *r.Name,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
