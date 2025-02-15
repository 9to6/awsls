// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListSsmParameter(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Ssmconn.DescribeParametersRequest(&ssm.DescribeParametersInput{})

	var result []terraform.Resource

	p := ssm.NewDescribeParametersPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Parameters {

			result = append(result, terraform.Resource{
				Type:      "aws_ssm_parameter",
				ID:        *r.Name,
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
