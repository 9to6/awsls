// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListDmsEndpoint(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Databasemigrationserviceconn.DescribeEndpointsRequest(&databasemigrationservice.DescribeEndpointsInput{})

	var result []terraform.Resource

	p := databasemigrationservice.NewDescribeEndpointsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Endpoints {

			result = append(result, terraform.Resource{
				Type:      "aws_dms_endpoint",
				ID:        *r.EndpointIdentifier,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
